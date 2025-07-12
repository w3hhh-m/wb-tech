package config

import (
	"time"

	"github.com/caarlos0/env/v11"
)

// Config describes main application configuration
type Config struct {
	// Hostname is a container name set by docker
	Hostname string `env:"HOSTNAME" envDefault:"localhost"`
	// LogLevel is a level of logs
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`

	// BrokerType is a type of broker used in application
	BrokerType string `env:"BROKER_TYPE,required,notEmpty"`
	// StorageType is a type of storage used in application
	StorageType string `env:"STORAGE_TYPE,required,notEmpty"`
	// CacheType is a type of cache used in application
	CacheType string `env:"CACHE_TYPE,required,notEmpty"`

	// ShutdownTimeout is a timeout for application graceful shutdown
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"10s"`
}

// LoadConfig loads application Config from environment variables.
// Returns error if something goes wrong while loading configuration
func LoadConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
