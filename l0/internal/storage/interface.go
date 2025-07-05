package storage

import "context"

// Storage interface
type Storage interface {
	Ping(ctx context.Context) error
}
