package broker

import "time"

// Broker interface
type Broker interface {
	// Close closes the Broker connection
	Close() error
	// Subscribe starts main subscription loop
	// and blocks until something goes wrong or
	// application is exiting. It takes handler which
	// will be called on every fetched message.
	// It must handle retries of message consumptions.
	// It takes MaxWorkers amount of messages and handles them concurrently.
	// Given handler must return error if something is wrong with actually message handling.
	// On handler error method will NOT commit message.
	// If something is wrong with the message itself (for example, invalid data)
	// handler must skip message and return nil to commit it
	Subscribe(handler func(message *Message) error)
}

// Message is a universal struct for all brokers messages.
// Other application packages will work with this type
type Message struct {
	// Key is a message key
	Key []byte
	// Value is a message value
	Value []byte
	// Timestamp is a message timestamp
	Timestamp time.Time
	// Headers is a message headers
	Headers map[string][]byte
}
