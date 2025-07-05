package local

import (
	"context"
	"fmt"
	"wb-tech-l0/internal/cache"
)

// Local is a Cache interface implementation for application in-memory cache
type Local struct {
	maxItems int
}

// New loads Local configuration and returns
// initialized Local implementation of Cache interface
func New() (cache.Cache, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading local cache config: %w", err)
	}
	_ = cfg
	return &Local{
		maxItems: cfg.MaxItems,
	}, nil
}

func (l *Local) Ping(ctx context.Context) error {
	// TODO: implement me
	return nil
}
