package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"wb-tech-l0/internal/broker"
	"wb-tech-l0/internal/broker/kafka"
	"wb-tech-l0/internal/cache"
	"wb-tech-l0/internal/cache/local"
	"wb-tech-l0/internal/config"
	"wb-tech-l0/internal/logger"
	zaplogger "wb-tech-l0/internal/logger/zap"
	"wb-tech-l0/internal/models"
	"wb-tech-l0/internal/registry"
	"wb-tech-l0/internal/storage"
	"wb-tech-l0/internal/storage/postgres"

	"github.com/go-playground/validator/v10"

	"golang.org/x/sync/errgroup"
)

// App is a struct that represents all application
type App struct {
	// cfg is application config
	cfg *config.Config
	// log is application logger
	log logger.Logger
	// ctx is application main context
	ctx context.Context

	// storage is a Storage client used in application
	storage storage.Storage
	// broker is a Broker client used in application
	broker broker.Broker
	// cache is a Cache client used in application
	cache cache.Cache

	// registries of supported services
	storageRegistry *registry.ServiceRegistry[storage.Storage]
	brokerRegistry  *registry.ServiceRegistry[broker.Broker]
	cacheRegistry   *registry.ServiceRegistry[cache.Cache]
	// we use registries to easily change the services used, even without changing the code.
	// when adding support for a new service, for example Redis for cache, we only need to register it
	// with a couple of lines of code. after that, we can choose which cache service to use (local or Redis)
	// at every application startup by simply changing the type in the .env configuration,
	// without changing a single line of code
}

// Start creates and runs the application with full lifecycle management.
// It loads the configuration, initializes the logger,
// creates tha application  and starts main application logic.
// After receiving exit signal it tries
// to gracefully shutdown application.
// It returns an error if something goes wrong.
func Start() error {
	// setting main application context cancelled on exit signals
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// loading application configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("could not load application config: %w", err)
	}

	// log implements Logger interface to avoid hardcoding exactly zap.
	// if we want to change log, we can just implement interface with new package
	// and change the line below
	log, err := zaplogger.New(cfg.LogLevel, cfg.Hostname)
	if err != nil {
		return fmt.Errorf("could not initialize log: %w", err)
	}

	log.Info("Successfully loaded application config")

	// creating new application instance
	app, err := New(ctx, cfg, log)
	if err != nil {
		return fmt.Errorf("could not create application: %w", err)
	}
	// flushing any buffered log entries before exiting
	defer app.log.Sync() // nolint: errcheck

	// running main application logic
	// this method will block until exit signal
	app.Run()

	// performing graceful shutdown
	app.Shutdown()
	return nil
}

// New creates and returns a new application instance.
// It registers supported services and creates clients for them
func New(ctx context.Context, cfg *config.Config, log logger.Logger) (*App, error) {
	app := &App{
		cfg: cfg,
		log: log,
		ctx: ctx,
		// creating registries of supported services.
		storageRegistry: registry.New[storage.Storage](),
		brokerRegistry:  registry.New[broker.Broker](),
		cacheRegistry:   registry.New[cache.Cache](),
	}

	// registering all supported services
	app.registerServices()

	// creating all service clients
	if err := app.createClients(); err != nil {
		// shutting down all connections if creating errors.
		// for example, if storage creating errors, but broker was created
		// we need to close broker
		app.Shutdown()
		return nil, fmt.Errorf("could not create clients: %w", err)
	}

	return app, nil
}

// Run runs the main application logic.
// It blocks until application exit signal.
func (a *App) Run() {
	a.log.Info("Application started successfully")

	// using single instance of validator for all messages
	// because it caches information about structs and validations
	validate := validator.New(validator.WithRequiredStructEnabled())

	// subscribe will block until something goes wrong or application is exiting.
	// given handler will be called on every successfully received message.
	// handler must return error if something is wrong with the message handling.
	// on error, broker will NOT commit message and there could be retries.
	a.broker.Subscribe(func(message *broker.Message) error {
		// add message key to log (this is handler's local logger)
		log := a.log.With(logger.Field("message_key", string(message.Key)))

		var order models.Order
		// parsing message value in order struct
		if err := json.Unmarshal(message.Value, &order); err != nil {
			log.Debug("Invalid JSON message. Handler skipping message", logger.Error(err))
			// returning nil to commit message in Subscribe
			return nil
		}

		// validating
		if err := validate.Struct(order); err != nil {
			log.Debug("Invalid order schema. Handler skipping message", logger.Error(err))
			// returning nil to commit message in Subscribe
			return nil
		}

		// adding order uid to logger for chaining storage logs with handler logs
		log = log.With(logger.Field("order_uid", order.OrderUID))

		// saving message
		err := a.storage.SaveOrder(&order)
		if err != nil {
			log.Warn("Failed to save order", logger.Error(err))
			if errors.Is(err, storage.ErrUniqueViolation) {
				log.Warn("Skipping order because it already exists")
				// returning nil to commit message in Subscribe because of invalid data
				return nil
			}
			// returning error to NOT commit message in broker
			return err
		}

		log.Debug("Order saved successfully")

		// TODO: cache

		return nil
	})

	a.log.Info("Got exiting signal. Shutting down application...", logger.Field("timeout", a.cfg.ShutdownTimeout))
}

// Shutdown performs graceful shutdown of all services.
// It tries to gracefully close all service connections within the timeout.
// All services are closing concurrently
func (a *App) Shutdown() {
	a.log.Info("Application shutting down gracefully")

	// creating shutdown context with timeout
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), a.cfg.ShutdownTimeout)
	defer shutdownCancel()

	// creating waitgroup for waiting connections closing.
	// we don't use errgroup because if any closing fails
	// we still need to try close other services
	var wg sync.WaitGroup

	// closing storage client
	if a.storage != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := a.storage.Close(); err != nil {
				a.log.Error("Could not close storage client", logger.Field("storage", a.cfg.StorageType), logger.Error(err))
				return
			}
			a.log.Info("Successfully closed storage client", logger.Field("storage", a.cfg.StorageType))
		}()
	}

	// closing broker client
	if a.broker != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := a.broker.Close(); err != nil {
				a.log.Error("Could not close broker client", logger.Field("broker", a.cfg.BrokerType), logger.Error(err))
				return
			}
			a.log.Info("Successfully closed broker client", logger.Field("broker", a.cfg.BrokerType))
		}()
	}

	// closing cache client
	if a.cache != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := a.cache.Close(); err != nil {
				a.log.Error("Could not close cache client", logger.Field("cache", a.cfg.CacheType), logger.Error(err))
				return
			}
			a.log.Info("Successfully closed cache client", logger.Field("cache", a.cfg.CacheType))
		}()
	}

	// done is closed when all services are closed
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	// waiting for all services being closed or shutdown timeout
	select {
	case <-done:
		a.log.Info("All services closed gracefully")
	case <-shutdownCtx.Done():
		a.log.Error("Shutdown timed out. Application exits gracelessly")
	}
}

// registerServices registers all supported services.
// To add support of new services, implement the required interface
// and register them inside this function.
func (a *App) registerServices() {
	a.storageRegistry.Register("postgres", func() (storage.Storage, error) {
		cfg, err := postgres.LoadConfig()
		if err != nil {
			return nil, fmt.Errorf("could not load postgres storage config: %w", err)
		}
		// add storage type to log
		return postgres.New(a.ctx, cfg, a.log.With(logger.Field("storage", "postgres")))
	})

	a.brokerRegistry.Register("kafka", func() (broker.Broker, error) {
		cfg, err := kafka.LoadConfig()
		if err != nil {
			return nil, fmt.Errorf("could not load kafka broker config: %w", err)
		}
		// add broker type to log
		return kafka.New(a.ctx, cfg, a.log.With(logger.Field("broker", "kafka")))
	})

	a.cacheRegistry.Register("local", func() (cache.Cache, error) {
		cfg, err := local.LoadConfig()
		if err != nil {
			return nil, fmt.Errorf("could not load local cache config: %w", err)
		}
		// add cache type to log
		return local.New(a.ctx, cfg, a.log.With(logger.Field("cache", "local")))
	})

	// you can add support of new services here by adding
	// them to registries as shown above
}

// createClients creates all service clients with provided types.
// All types must be registered in registries (in registerServices function)
// and services required configuration must be set in environment.
// It creates all service clients concurrently.
func (a *App) createClients() error {
	// creating errgroup to wait for creations and check errors
	g, _ := errgroup.WithContext(a.ctx)

	// creating storageClient concurrently
	g.Go(func() error {
		// creating storage client with provided StorageType.
		// StorageType must be registered in storageRegistry (in registerServices function)
		// and its required configuration must be set in environment.
		storageClient, err := a.storageRegistry.Create(a.cfg.StorageType)
		if err != nil {
			return fmt.Errorf("could not create storage client: %w", err)
		}
		a.storage = storageClient
		a.log.Info("Successfully created storage client", logger.Field("storage", a.cfg.StorageType))
		return nil
	})

	// creating brokerClient concurrently
	g.Go(func() error {
		// creating broker client with provided BrokerType.
		// BrokerType must be registered in brokerRegistry (in registerServices function)
		// and its required configuration must be set in environment.
		brokerClient, err := a.brokerRegistry.Create(a.cfg.BrokerType)
		if err != nil {
			return fmt.Errorf("could not create broker client: %w", err)
		}
		a.broker = brokerClient
		a.log.Info("Successfully created broker client", logger.Field("broker", a.cfg.BrokerType))
		return nil
	})

	// creating cacheClient concurrently
	g.Go(func() error {
		// creating cache client with provided CacheType.
		// CacheType must be registered in cacheRegistry (in registerServices function)
		// and its required configuration must be set in environment.
		cacheClient, err := a.cacheRegistry.Create(a.cfg.CacheType)
		if err != nil {
			return fmt.Errorf("could not create cache client: %w", err)
		}
		a.cache = cacheClient
		a.log.Info("Successfully created cache client", logger.Field("cache", a.cfg.CacheType))
		return nil
	})

	// waiting for all creations
	if err := g.Wait(); err != nil {
		// if any creation fails, return the error
		return err
		// if createClients errors, then Shutdown will be called
		// and all already created clients will be closed
	}

	return nil
}
