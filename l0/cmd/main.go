package main

import (
	"fmt"
	"os"
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

func main() {
	// TODO: log.Fatal will exit, and `defer log.Sync()` will not run
	// TODO: move logic to some function, not main

	// TODO: update README

	// loading application configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Could not load application config: %s\n", err)
		os.Exit(1)
	}

	// log implements Logger interface to avoid hardcoding exactly zap.
	// if we want to change log, we can just implement interface with new package
	// and change the line below
	log, err := zaplogger.New(cfg.LogLevel, cfg.Hostname)
	if err != nil {
		fmt.Printf("Could not initialize log: %s\n", err)
		os.Exit(1)
	}
	// flushing any buffered log entries before exiting
	defer log.Sync() // nolint: errcheck

	log.Info("Successfully loaded application config")

	// creating registries of supported services.
	storageRegistry := registry.New[storage.Storage]()
	brokerRegistry := registry.New[broker.Broker]()
	cacheRegistry := registry.New[cache.Cache]()

	// we use registries to easily change the services used, even without changing the code.
	// when adding support for a new service, for example Redis for cache, we only need to register it
	// with a couple of lines of code. after that, we can choose which cache service to use (local or Redis)
	// by simply changing the type in the .env configuration, without changing a single line of code

	// registration of supported services
	storageRegistry.Register("postgres", func() (storage.Storage, error) {
		cfg, err := postgres.LoadConfig()
		if err != nil {
			return nil, fmt.Errorf("could not load postgres storage config: %w", err)
		}
		return postgres.New(cfg, log)
	})
	brokerRegistry.Register("kafka", func() (broker.Broker, error) {
		cfg, err := kafka.LoadConfig()
		if err != nil {
			return nil, fmt.Errorf("could not load kafka broker config: %w", err)
		}
		return kafka.New(cfg, log)
	})
	cacheRegistry.Register("local", func() (cache.Cache, error) {
		cfg, err := local.LoadConfig()
		if err != nil {
			return nil, fmt.Errorf("could not load local cache config: %w", err)
		}
		return local.New(cfg, log)
	})
	// to add support for new services, implement the required interface
	// and register them as shown above

	// creating storageClient with provided StorageType.
	// StorageType must be registered in storageRegistry
	// and its required configuration must be set in environment.
	storageClient, err := storageRegistry.Create(cfg.StorageType)
	if err != nil {
		log.Fatal("Could not create storage client", logger.Error(err))
	}
	log.Info("Successfully created storage client", logger.Field("type", cfg.StorageType))

	// creating brokerClient with provided BrokerType.
	// BrokerType must be registered in brokerRegistry
	// and its required configuration must be set in environment.
	brokerClient, err := brokerRegistry.Create(cfg.BrokerType)
	if err != nil {
		log.Fatal("Could not create broker client", logger.Error(err))
	}
	log.Info("Successfully created broker client", logger.Field("type", cfg.BrokerType))

	// creating cacheClient with provided CacheType.
	// CacheType must be registered in cacheRegistry
	// and its required configuration must be set in environment.
	cacheClient, err := cacheRegistry.Create(cfg.CacheType)
	if err != nil {
		log.Fatal("Could not create cache client", logger.Error(err))
	}
	log.Info("Successfully created cache client", logger.Field("type", cfg.CacheType))

	_, _, _ = storageClient, brokerClient, cacheClient
}
