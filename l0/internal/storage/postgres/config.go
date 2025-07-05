package postgres

import (
	"github.com/caarlos0/env/v11"
	"time"
)

// Config describes Postgres storage configuration
type Config struct {
	Host           string        `env:"POSTGRES_HOST,required,notEmpty"`
	Port           string        `env:"POSTGRES_PORT,required,notEmpty"`
	User           string        `env:"POSTGRES_USER,required,notEmpty"`
	Password       string        `env:"POSTGRES_PASSWORD,required,notEmpty"`
	Database       string        `env:"POSTGRES_DATABASE,required,notEmpty"`
	SSLMode        string        `env:"POSTGRES_SSLMODE" envDefault:"disable"`
	ConnectTimeout time.Duration `env:"POSTGRES_CONNECT_TIMEOUT" envDefault:"5s"`

	PoolMaxConns    int32         `env:"POSTGRES_POOL_MAX_CONNS" envDefault:"10"`
	PoolMinConns    int32         `env:"POSTGRES_POOL_MIN_CONNS" envDefault:"2"`
	MaxConnLifetime time.Duration `env:"POSTGRES_MAX_CONN_LIFETIME" envDefault:"1h"`
	MaxConnIdleTime time.Duration `env:"POSTGRES_MAX_CONN_IDLE_TIME" envDefault:"30m"`

	RequestTimeout time.Duration `env:"POSTGRES_REQUEST_TIMEOUT" envDefault:"5s"`
	RetryTimeout   time.Duration `env:"POSTGRES_RETRY_TIMEOUT" envDefault:"5s"`
	MaxRetries     int           `env:"POSTGRES_MAX_RETRIES" envDefault:"3"`
}

// LoadConfig loads Postgres storage Config from environment variables.
// Returns error if something goes wrong while loading configuration
func LoadConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
