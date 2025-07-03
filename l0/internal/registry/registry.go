package registry

import (
	"fmt"
	"sync"
)

// ServiceRegistry is generic concurrent-safe registry that maps services names to constructor functions
// T specifies the service interface type that implementations must satisfy
type ServiceRegistry[T interface{}] struct {
	mu        sync.RWMutex
	functions map[string]func() (T, error)
}

// New creates and returns empty ServiceRegistry for type T
func New[T interface{}]() *ServiceRegistry[T] {
	return &ServiceRegistry[T]{
		functions: make(map[string]func() (T, error)),
	}
}

// Register adds to the ServiceRegistry new constructor function of service with provided name.
// This method is safe for concurrent use
func (r *ServiceRegistry[T]) Register(name string, function func() (T, error)) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.functions[name] = function
}

// Create creates a new service instance of type T using the constructor registered with provided name.
// Returns an error if no constructor is registered for the name.
// This method is safe for concurrent use
func (r *ServiceRegistry[T]) Create(name string) (T, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	fn, ok := r.functions[name]
	if !ok {
		var zero T
		return zero, fmt.Errorf("unknown type: %s", name)
	}
	return fn()
}
