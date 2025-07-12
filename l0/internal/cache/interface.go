package cache

// Cache interface
type Cache interface {
	// Close closes the Cache connection
	Close() error
}
