package local

import (
	"github.com/caarlos0/env/v11"
)

// Config describes Local cache configuration
type Config struct {
	MaxItems int `env:"LOCAL_CACHE_MAX_ITEMS" envDefault:"1000"`
}

// LoadConfig loads Local cache Config from environment variables.
// Returns error if something goes wrong while loading configuration
func LoadConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
