package local

import (
	"context"
	"sync"
	"time"

	"wb-tech-l0/internal/logger"
)

// Local is a Cache interface implementation for application in-memory cache.
// It's methods are safe for concurrent use
type Local struct {
	maxItems int
	ttl      time.Duration

	items map[string]cacheItem
	mu    sync.RWMutex

	ctx context.Context
	log logger.Logger
}

type cacheItem struct {
	value     interface{}
	expiresAt time.Time
}

// New creates and returns initialized Local implementation of Cache interface
func New(ctx context.Context, cfg *Config, log logger.Logger) (*Local, error) {
	log.Debug("Creating cache connection")
	return &Local{
		maxItems: cfg.MaxItems,
		ttl:      cfg.TTL,
		items:    make(map[string]cacheItem),
		log:      log,
		ctx:      ctx,
	}, nil
}

// Close closes the Local cache connection
func (l *Local) Close() error {
	// nothing to close for in-memory cache
	return nil
}

// GetOrder gets order from cache if exists and not expired.
// It also handles lazy deletion of getting expired keys
func (l *Local) GetOrder(key string) (interface{}, bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	l.log.Debug("Attempting to get order", logger.Field("key", key))

	item, found := l.items[key]
	if !found {
		return nil, false
	}

	// check if item expired
	if time.Now().After(item.expiresAt) {
		// lazy ttl removing
		l.log.Debug("Lazy cache expired item removal", logger.Field("key", key))
		delete(l.items, key)
		return nil, false
	}

	return item.value, true
}

// SaveOrder saves order to cache
// It also handles the removing expired keys or random one
// if reached the maximum capacity
func (l *Local) SaveOrder(key string, value interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.log.Debug("Attempting to save order", logger.Field("key", key))

	// clean expired items if reaching max capacity
	if len(l.items) >= l.maxItems {
		l.log.Debug("Reached cache maximum capacity. Cleaning all expired keys")
		l.cleanExpired()
		// if still there is no space, remove random item
		if len(l.items) >= l.maxItems {
			// range over map is random
			for k := range l.items {
				l.log.Debug("Reached cache maximum capacity. Removing random item", logger.Field("key", k))
				delete(l.items, k)
				break
			}
		}
	}

	l.items[key] = cacheItem{
		value:     value,
		expiresAt: time.Now().Add(l.ttl),
	}
}

// cleanExpired removes all expired items from cache
func (l *Local) cleanExpired() {
	now := time.Now()
	for k, v := range l.items {
		if now.After(v.expiresAt) {
			l.log.Debug("Active cache expired item removal", logger.Field("key", k))
			delete(l.items, k)
		}
	}
}
