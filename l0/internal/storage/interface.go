package storage

import "wb-tech-l0/internal/models"

// Storage interface
type Storage interface {
	// Close closes the Storage connection
	Close() error
	// SaveOrder takes order and saves it to storage.
	// It also must handle the retries of saving
	SaveOrder(order *models.Order) error
}
