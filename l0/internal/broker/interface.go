package broker

import "context"

// Broker interface
type Broker interface {
	Ping(ctx context.Context) error
}
