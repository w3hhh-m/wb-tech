package postgres

import (
	"time"

	"github.com/caarlos0/env/v11"
)

// Config describes Postgres storage configuration
type Config struct {
	// single connection configuration
	// Host is a Postgres host to connect to.
	Host string `env:"POSTGRES_HOST,required,notEmpty"`
	// Port is a Postgres port to connect to.
	Port string `env:"POSTGRES_PORT,required,notEmpty"`
	// User is a Postgres user to connect with.
	User string `env:"POSTGRES_USER,required,notEmpty"`
	// Password is a Postgres password to connect with.
	Password string `env:"POSTGRES_PASSWORD,required,notEmpty"`
	// Database is a Postgres database to connect to.
	Database string `env:"POSTGRES_DB,required,notEmpty"`
	// SSLMode is a Postgres SSL mode to connect with.
	SSLMode string `env:"POSTGRES_SSLMODE" envDefault:"disable"`
	// ConnectTimeout is a timeout for connecting to Postgres.
	ConnectTimeout time.Duration `env:"POSTGRES_CONNECT_TIMEOUT" envDefault:"5s"`

	// pool configuration
	// PoolMaxConns is a maximum number of connections in the pool.
	PoolMaxConns int32 `env:"POSTGRES_POOL_MAX_CONNS" envDefault:"10"`
	// PoolMinConns is a minimum number of connections in the pool.
	PoolMinConns int32 `env:"POSTGRES_POOL_MIN_CONNS" envDefault:"2"`
	// MaxConnLifetime is a maximum lifetime of a connection in the pool.
	MaxConnLifetime time.Duration `env:"POSTGRES_MAX_CONN_LIFETIME" envDefault:"1h"`
	// MaxConnIdleTime is a maximum idle time of a connection in the pool.
	MaxConnIdleTime time.Duration `env:"POSTGRES_MAX_CONN_IDLE_TIME" envDefault:"30m"`

	// custom retry configuration
	// RequestTimeout is a timeout for request to Postgres.
	RequestTimeout time.Duration `env:"POSTGRES_REQUEST_TIMEOUT" envDefault:"5s"`
	// RetryTimeout is a timeout for retrying operations.
	RetryTimeout time.Duration `env:"POSTGRES_RETRY_TIMEOUT" envDefault:"5s"`
	// MaxRetries is a maximum number of retries for operations.
	MaxRetries int `env:"POSTGRES_MAX_RETRIES" envDefault:"3"`
}

// LoadConfig loads Postgres storage Config from environment variables.
// Returns error if something goes wrong while loading configuration
func LoadConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	// not validating required fields here, because it`s validated while parsing
	if cfg.PoolMinConns < 0 {
		cfg.PoolMinConns = 2
	}
	if cfg.PoolMaxConns < 0 {
		cfg.PoolMaxConns = 10
	}
	if cfg.MaxRetries < 0 {
		cfg.MaxRetries = 3
	}

	return cfg, nil
}
