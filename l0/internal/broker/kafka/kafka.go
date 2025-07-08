package kafka

import (
	"context"
	kafkago "github.com/segmentio/kafka-go"
	"time"
	"wb-tech-l0/internal/logger"
)

// Kafka is a Broker interface implementation for Kafka
type Kafka struct {
	reader       *kafkago.Reader
	readTimeout  time.Duration
	retryTimeout time.Duration
	maxRetries   int

	logger logger.Logger
}

// New creates and returns initialized Kafka implementation of Storage interface
func New(cfg *Config, logger logger.Logger) (*Kafka, error) {
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
		logger:       logger,
	}, nil
}

func (k *Kafka) Ping(ctx context.Context) error {
	// TODO: implement me
	return nil
}
