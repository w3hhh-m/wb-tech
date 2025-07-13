package broker

// Broker interface
type Broker interface {
	// Close closes the Broker connection
	Close() error
	// Subscribe starts main subscription loop
	// and blocks until something goes wrong or
	// application is exiting. It takes handler which
	// will be called on every fetched message
	Subscribe(handler func(message []byte) error)
}
