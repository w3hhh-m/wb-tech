package storage

// Storage interface
type Storage interface {
	// Close closes the Storage connection
	Close() error
}
