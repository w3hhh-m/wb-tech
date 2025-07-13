package local

import (
	"context"
	"wb-tech-l0/internal/logger"
)

// Local is a Cache interface implementation for application in-memory cache
type Local struct {
	maxItems int

	ctx context.Context
	log logger.Logger
}

// New creates and returns initialized Local implementation of Cache interface
func New(ctx context.Context, cfg *Config, log logger.Logger) (*Local, error) {
	log.Debug("Creating cache connection")
	return &Local{
		maxItems: cfg.MaxItems,
		log:      log,
		ctx:      ctx,
	}, nil
}

// Close closes the Local cache connection
func (l *Local) Close() error {
	// nothing to close for in-memory cache
	return nil
}
