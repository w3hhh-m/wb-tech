package postgres

import (
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
)

// Config describes Postgres storage configuration
type Config struct {
	// single connection configuration
	// Host is a Postgres host to connect to.
	Host string `env:"POSTGRES_HOST,required,notEmpty"`
	// Port is a Postgres port to connect to.
	Port string `env:"POSTGRES_PORT,required,notEmpty" validate:"numeric"`
	// User is a Postgres user to connect with.
	User string `env:"POSTGRES_USER,required,notEmpty"`
	// Password is a Postgres password to connect with.
	Password string `env:"POSTGRES_PASSWORD,required,notEmpty"`
	// Database is a Postgres database to connect to.
	Database string `env:"POSTGRES_DB,required,notEmpty"`
	// SSLMode is a Postgres SSL mode to connect with.
	SSLMode string `env:"POSTGRES_SSLMODE" envDefault:"disable" validate:"oneof=disable require verify-ca verify-full"`
	// ConnectTimeout is a timeout for connecting to Postgres.
	ConnectTimeout time.Duration `env:"POSTGRES_CONNECT_TIMEOUT" envDefault:"5s" validate:"gte=100ms"`

	// pool configuration
	// PoolMaxConns is a maximum number of connections in the pool.
	PoolMaxConns int32 `env:"POSTGRES_POOL_MAX_CONNS" envDefault:"10" validate:"gte=1,gtefield=PoolMinConns"`
	// PoolMinConns is a minimum number of connections in the pool.
	PoolMinConns int32 `env:"POSTGRES_POOL_MIN_CONNS" envDefault:"2" validate:"gte=1,ltefield=PoolMaxConns"`
	// MaxConnLifetime is a maximum lifetime of a connection in the pool.
	MaxConnLifetime time.Duration `env:"POSTGRES_MAX_CONN_LIFETIME" envDefault:"1h" validate:"gte=1s"`
	// MaxConnIdleTime is a maximum idle time of a connection in the pool.
	MaxConnIdleTime time.Duration `env:"POSTGRES_MAX_CONN_IDLE_TIME" envDefault:"30m" validate:"gte=1s"`

	// custom retry configuration
	// RequestTimeout is a timeout for request to Postgres.
	RequestTimeout time.Duration `env:"POSTGRES_REQUEST_TIMEOUT" envDefault:"5s" validate:"gte=100ms"`
	// RetryTimeout is a timeout for retrying operations.
	RetryTimeout time.Duration `env:"POSTGRES_RETRY_TIMEOUT" envDefault:"5s" validate:"gte=100ms"`
	// MaxRetries is a maximum number of retries for operations.
	MaxRetries int `env:"POSTGRES_MAX_RETRIES" envDefault:"3" validate:"gte=0"`
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
	validate := validator.New()
	err = validate.Struct(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
