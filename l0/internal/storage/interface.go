package storage

import (
	"context"
	"wb-tech-l0/internal/models"
)

// Storage interface
type Storage interface {
	// Close closes the Storage connection
	Close() error
	// SaveOrder takes order and saves it to storage.
	// It also must handle the retries of saving
	SaveOrder(order *models.Order) error
	// GetOrder takes user request context and order uid and fetches its model.
	// It also must handle the retries of fetching
	GetOrder(ctx context.Context, uid string) (*models.Order, error)
}
