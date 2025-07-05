package kafka

import (
	"context"
	"fmt"
	kafkago "github.com/segmentio/kafka-go"
	"time"
	"wb-tech-l0/internal/broker"
)

// Kafka is a Broker interface implementation for Kafka
type Kafka struct {
	reader       *kafkago.Reader
	readTimeout  time.Duration
	retryTimeout time.Duration
	maxRetries   int
}

// New loads Kafka configuration and
// returns initialized Kafka implementation of Storage interface
func New() (broker.Broker, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading kafka broker config: %w", err)
	}

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
	}, nil
}

func (k *Kafka) Ping(ctx context.Context) error {
	// TODO: implement me
	return nil
}
