package cache

// Cache interface
type Cache interface {
	// Close closes the Cache connection
	Close() error
	// GetOrder gets order from cache
	GetOrder(key string) (interface{}, bool)
	// SaveOrder saves order to cache
	SaveOrder(key string, value interface{})
}
