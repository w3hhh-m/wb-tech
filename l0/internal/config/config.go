package config

import "github.com/caarlos0/env/v11"

// Config describes main application configuration
type Config struct {
	LogLevel    string `env:"LOG_LEVEL" envDefault:"info"`
	BrokerType  string `env:"BROKER_TYPE,required,notEmpty"`
	StorageType string `env:"STORAGE_TYPE,required,notEmpty"`
	CacheType   string `env:"CACHE_TYPE,required,notEmpty"`
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
