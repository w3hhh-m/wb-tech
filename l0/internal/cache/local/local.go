package local

import (
	"context"
	"wb-tech-l0/internal/logger"
)

// Local is a Cache interface implementation for application in-memory cache
type Local struct {
	maxItems int

	logger logger.Logger
}

// New creates and returns initialized Local implementation of Cache interface
func New(cfg *Config, logger logger.Logger) (*Local, error) {
	return &Local{
		maxItems: cfg.MaxItems,
		logger:   logger,
	}, nil
}

func (l *Local) Ping(ctx context.Context) error {
	// TODO: implement me
	return nil
}
