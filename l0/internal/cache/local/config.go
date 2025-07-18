package local

import (
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
)

// Config describes Local cache configuration
type Config struct {
	// MaxItems is a maximum number of items that can be stored in Local cache
	MaxItems int `env:"LOCAL_CACHE_MAX_ITEMS" envDefault:"1000" validate:"gte=1"`
	// TTL is time-to-live for cache items
	TTL time.Duration `env:"LOCAL_CACHE_TTL" envDefault:"3600s" validate:"gte=1s"`

	// no retries on Local cache operations
}

// LoadConfig loads Local cache Config from environment variables.
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
