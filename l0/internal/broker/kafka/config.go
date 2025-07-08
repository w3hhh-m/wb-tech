package kafka

import (
	"time"

	"github.com/caarlos0/env/v11"
)

// Config describes Kafka broker configuration
type Config struct {
	Brokers     []string `env:"KAFKA_BROKERS,required,notEmpty" envSeparator:","`
	Topic       string   `env:"KAFKA_TOPIC,required,notEmpty"`
	GroupID     string   `env:"KAFKA_GROUP_ID,required,notEmpty"`
	StartOffset int64    `env:"KAFKA_START_OFFSET" envDefault:"-2"`
	MinBytes    int      `env:"KAFKA_MIN_BYTES" envDefault:"1"`
	MaxBytes    int      `env:"KAFKA_MAX_BYTES" envDefault:"10e6"`

	ReadTimeOut  time.Duration `env:"KAFKA_READ_TIMEOUT" envDefault:"5s"`
	RetryTimeOut time.Duration `env:"KAFKA_RETRY_TIMEOUT" envDefault:"5s"`
	MaxRetries   int           `env:"KAFKA_MAX_RETRIES" envDefault:"3"`
}

// LoadConfig loads Kafka broker Config from environment variables.
// Returns error if something goes wrong while loading configuration
func LoadConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
