package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
	"wb-tech-l0/internal/logger"
)

// Postgres is a Storage interface implementation for PostgreSQL
type Postgres struct {
	pool           *pgxpool.Pool
	requestTimeout time.Duration
	retryTimeout   time.Duration
	maxRetries     int

	logger logger.Logger
}

// New creates and returns initialized PostgreSQL implementation of Storage interface
func New(cfg *Config, logger logger.Logger) (*Postgres, error) {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode,
	)

	dbpoolCfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("error parsing postgres storage config: %w", err)
	}

	dbpoolCfg.ConnConfig.ConnectTimeout = cfg.ConnectTimeout
	dbpoolCfg.MaxConns = cfg.PoolMaxConns
	dbpoolCfg.MinConns = cfg.PoolMinConns
	dbpoolCfg.MaxConnLifetime = cfg.MaxConnLifetime
	dbpoolCfg.MaxConnIdleTime = cfg.MaxConnIdleTime

	pool, err := pgxpool.NewWithConfig(context.Background(), dbpoolCfg)
	if err != nil {
		return nil, fmt.Errorf("error connecting to postgres storage: %w", err)
	}

	return &Postgres{
		pool:           pool,
		requestTimeout: cfg.RequestTimeout,
		retryTimeout:   cfg.RetryTimeout,
		maxRetries:     cfg.MaxRetries,
		logger:         logger,
	}, nil
}

func (p *Postgres) Ping(ctx context.Context) error {
	// TODO: implement me
	return nil
}
