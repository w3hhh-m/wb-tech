package cache

import "context"

// Cache interface
type Cache interface {
	Ping(ctx context.Context) error
}
