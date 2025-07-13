package kafka

import (
	"context"
	"time"
	"wb-tech-l0/internal/logger"

	kafkago "github.com/segmentio/kafka-go"
)

// Kafka is a Broker interface implementation for Kafka
type Kafka struct {
	reader       *kafkago.Reader
	readTimeout  time.Duration
	retryTimeout time.Duration
	maxRetries   int

	ctx context.Context
	log logger.Logger
}

// New creates and returns initialized Kafka implementation of Broker interface
func New(ctx context.Context, cfg *Config, log logger.Logger) (*Kafka, error) {
	log.Debug("Creating broker connection")

	kafkaCfg := kafkago.ReaderConfig{
		Brokers:          cfg.Brokers,
		Topic:            cfg.Topic,
		GroupID:          cfg.GroupID,
		StartOffset:      cfg.StartOffset,
		MinBytes:         cfg.MinBytes,
		MaxBytes:         cfg.MaxBytes,
		ReadBatchTimeout: cfg.ReadTimeOut,
		Logger:           nil,
		ErrorLogger:      nil,
		MaxAttempts:      cfg.MaxRetries,
	}

	reader := kafkago.NewReader(kafkaCfg)

	return &Kafka{
		reader:       reader,
		readTimeout:  cfg.ReadTimeOut,
		retryTimeout: cfg.RetryTimeOut,
		maxRetries:   cfg.MaxRetries,
		log:          log,
		ctx:          ctx,
	}, nil
}

// Close closes the Kafka broker connection
func (k *Kafka) Close() error {
	return k.reader.Close()
}

// Subscribe starts main Kafka broker subscription loop
// and blocks until something goes wrong or application is exiting.
// It takes handler which will be called on every fetched message
func (k *Kafka) Subscribe(handler func(message []byte) error) {
	// add stats to log
	stats := k.reader.Stats()
	log := k.log.With(logger.Field("client_id", stats.ClientID), logger.Field("topic", stats.Topic))

	log.Debug("Starting broker subscription loop")
	defer log.Debug("Broker subscription loop exited")

	// main loop
	for {
		// select on ctx to return when application is exiting
		select {
		case <-k.ctx.Done():
			log.Debug("Context cancelled during subscription loop")
			return
		default:
		}
		// fetching message. this call will block until
		// we got message or error or context is cancelled
		msg, err := k.reader.FetchMessage(k.ctx)
		if err != nil {
			// if error is about context cancelling
			if k.ctx.Err() != nil {
				log.Debug("Context cancelled during message reading")
				return
			}

			// error is not about context
			log.Warn("Error fetching broker message", logger.Error(err))

			// waiting retry timeout and try to fetch message again
			// select on ctx to return when application is exiting
			select {
			case <-k.ctx.Done():
				log.Debug("Context cancelled during retrying message reading")
				return
			case <-time.After(k.retryTimeout):
				// continue fetching
			}
			// to the start of for loop to fetch message again
			continue
		}

		// add message key to log
		log = log.With(logger.Field("message_key", string(msg.Key)))
		// logging message value
		log.Debug("Message received", logger.Field("message_value", string(msg.Value)))

		// now when we got message we need to handle it
		// retries of handling must be handled in handler
		if err = handler(msg.Value); err != nil {
			log.Warn("Message handler returned error. Skipping message", logger.Error(err))
			// NOT COMMITING MESSAGE ON HANDLER ERROR
			continue
		}

		// COMMIT ONLY IF MESSAGE HANDLED SUCCESSFULLY
		if err = k.reader.CommitMessages(k.ctx, msg); err != nil {
			log.Warn("Failed to commit broker message", logger.Error(err))
		}
	}
}
