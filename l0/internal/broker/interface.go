package broker

// Broker interface
type Broker interface {
	// Close closes the Broker connection
	Close() error
}
