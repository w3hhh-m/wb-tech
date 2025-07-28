package kafka

import (
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
)

// Config describes Kafka broker configuration
type Config struct {
	// Brokers is a list of Kafka brokers to connect to.
	Brokers []string `env:"KAFKA_BROKERS,required,notEmpty" envSeparator:","`
	// Topic is a Kafka topic to consume messages from.
	Topic string `env:"KAFKA_TOPIC,required,notEmpty"`
	// GroupID is a Kafka group ID to consume messages from.
	GroupID string `env:"KAFKA_GROUP_ID,required,notEmpty"`
	// StartOffset is a Kafka start offset to consume messages from.
	// -2 means to consume from the least recent offset.
	// -1 means to consume from the most recent offset.
	StartOffset int64 `env:"KAFKA_START_OFFSET" envDefault:"-2" validate:"oneof=-2 -1"`
	// MinBytes is a minimum number of bytes to read from Kafka.
	MinBytes int `env:"KAFKA_MIN_BYTES" envDefault:"1" validate:"gte=1,ltefield=MaxBytes"`
	// MaxBytes is a maximum number of bytes to read from Kafka.
	MaxBytes int `env:"KAFKA_MAX_BYTES" envDefault:"10e6" validate:"gte=1,gtefield=MinBytes"`
	// ReadTimeOut is a timeout for reading from Kafka.
	ReadTimeOut time.Duration `env:"KAFKA_READ_TIMEOUT" envDefault:"5s" validate:"gte=100ms"`

	// MaxWorkers is a maximum number of concurrent workers for messages handling
	MaxWorkers int `env:"MAX_WORKERS" envDefault:"1" validate:"gte=1"`

	// custom retry configuration
	// RetryTimeOut is a timeout for retrying operations.
	RetryTimeOut time.Duration `env:"KAFKA_RETRY_TIMEOUT" envDefault:"5s" validate:"gte=100ms"`
	// MaxRetries is a maximum number of retries for operations.
	MaxRetries int `env:"KAFKA_MAX_RETRIES" envDefault:"3" validate:"gte=0"`
}

// LoadConfig loads Kafka broker Config from environment variables.
// Returns error if something goes wrong while loading configuration
func LoadConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	// not validating required fields here, because it`s validated while parsing
	validate := validator.New()
	err = validate.Struct(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
