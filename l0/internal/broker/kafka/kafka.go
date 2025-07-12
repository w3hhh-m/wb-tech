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
	log.Debug("Attempting to create broker connection")

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
