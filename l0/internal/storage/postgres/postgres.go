package postgres

import (
	"context"
	"fmt"
	"time"
	"wb-tech-l0/internal/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Postgres is a Storage interface implementation for PostgreSQL
type Postgres struct {
	pool           *pgxpool.Pool
	requestTimeout time.Duration
	retryTimeout   time.Duration
	maxRetries     int

	ctx context.Context
	log logger.Logger
}

// New creates and returns initialized Postgres implementation of Storage interface
func New(ctx context.Context, cfg *Config, log logger.Logger) (*Postgres, error) {
	log.Debug("Creating storage connection")

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

	pool, err := pgxpool.NewWithConfig(ctx, dbpoolCfg)
	if err != nil {
		return nil, fmt.Errorf("error connecting to postgres storage: %w", err)
	}

	return &Postgres{
		pool:           pool,
		requestTimeout: cfg.RequestTimeout,
		retryTimeout:   cfg.RetryTimeout,
		maxRetries:     cfg.MaxRetries,
		log:            log,
		ctx:            ctx,
	}, nil
}

// Close closes the Postgres storage connection
func (p *Postgres) Close() error {
	p.pool.Close()
	return nil
}
