package app

import (
	"context"
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
	"wb-tech-l0/internal/registry"
	"wb-tech-l0/internal/storage"
	"wb-tech-l0/internal/storage/postgres"
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

// RunApp creates and runs the application with full lifecycle management.
// It loads the configuration, initializes the logger,
// registers services, creates the clients,
// and starts main application logic.
// After receiving exit signal it tries
// to gracefully shutdown application.
// It returns an error if something goes wrong.
func RunApp() error {
	// setting main application context cancelled on exit signals
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// creating new application instance
	app, err := New(ctx)
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
func New(ctx context.Context) (*App, error) {
	// loading application configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("could not load application config: %w", err)
	}

	// log implements Logger interface to avoid hardcoding exactly zap.
	// if we want to change log, we can just implement interface with new package
	// and change the line below
	log, err := zaplogger.New(cfg.LogLevel, cfg.Hostname)
	if err != nil {
		return nil, fmt.Errorf("could not initialize log: %w", err)
	}

	log.Info("Successfully loaded application config")

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
	if err = app.createClients(); err != nil {
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

	// TODO: main logic

	// waiting for application exit signal
	<-a.ctx.Done()
	a.log.Info("Got exiting signal. Shutting down application...", logger.Field("timeout", a.cfg.ShutdownTimeout))
}

// Shutdown performs graceful shutdown of all services.
// It tries to gracefully close all service connections within the timeout.
// All services Close methods are called concurrently
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
func (a *App) createClients() error {
	// creating storage client with provided StorageType.
	// StorageType must be registered in storageRegistry (in registerServices function)
	// and its required configuration must be set in environment.
	storageClient, err := a.storageRegistry.Create(a.cfg.StorageType)
	if err != nil {
		return fmt.Errorf("could not create storage client: %w", err)
	}
	a.storage = storageClient
	a.log.Info("Successfully created storage client", logger.Field("storage", a.cfg.StorageType))

	// creating broker client with provided BrokerType.
	// BrokerType must be registered in brokerRegistry (in registerServices function)
	// and its required configuration must be set in environment.
	brokerClient, err := a.brokerRegistry.Create(a.cfg.BrokerType)
	if err != nil {
		return fmt.Errorf("could not create broker client: %w", err)
	}
	a.broker = brokerClient
	a.log.Info("Successfully created broker client", logger.Field("broker", a.cfg.BrokerType))

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
}
