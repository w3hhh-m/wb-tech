package config

import (
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
)

// Config describes main application configuration
type Config struct {
	// Hostname is a container name set by docker
	Hostname string `env:"HOSTNAME" envDefault:"localhost"`
	// LogLevel is a level of logs
	LogLevel string `env:"LOG_LEVEL" envDefault:"info" validate:"oneof=debug info warn error"`

	// BrokerType is a type of broker used in application
	BrokerType string `env:"BROKER_TYPE,required,notEmpty"`
	// StorageType is a type of storage used in application
	StorageType string `env:"STORAGE_TYPE,required,notEmpty"`
	// CacheType is a type of cache used in application
	CacheType string `env:"CACHE_TYPE,required,notEmpty"`

	// Server is the HTTP server configuration
	Server ServerConfig
	// ShutdownTimeout is a timeout for application graceful shutdown
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"10s" validate:"gte=1s"`
}

// ServerConfig describes HTTP server configuration
// It is part of main application config, so I declared it here
// (unlike configs for external services, that can be switched)
type ServerConfig struct {
	// Address is the address the HTTP server listens on
	Address string `env:"HTTP_ADDRESS" envDefault:":8080"`
	// ReadTimeout is the maximum duration for reading the entire request
	ReadTimeout time.Duration `env:"HTTP_READ_TIMEOUT" envDefault:"5s" validate:"gte=1s"`
	// WriteTimeout is the maximum duration before timing out writes of the response
	WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT" envDefault:"10s" validate:"gte=1s"`
	// IdleTimeout is the maximum amount of time to wait for the next request
	IdleTimeout time.Duration `env:"HTTP_IDLE_TIMEOUT" envDefault:"120s" validate:"gte=1s"`
}

// LoadConfig loads application Config from environment variables.
// Returns error if something goes wrong while loading configuration
func LoadConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	validate := validator.New()
	err = validate.Struct(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
