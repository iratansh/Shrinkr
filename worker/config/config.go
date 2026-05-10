package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds runtime configuration for the analytics worker.
type Config struct {
	PostgresDSN     string
	SQSQueueURL     string
	SQSEndpoint     string
	AWSRegion       string
	BatchSize       int32
	WaitTimeSeconds int32
}

// Load reads configuration from environment variables.
func Load() (*Config, error) {
	batchSize, err := parseInt32("SQS_BATCH_SIZE", 10)
	if err != nil {
		return nil, err
	}

	waitTimeSeconds, err := parseInt32("SQS_WAIT_TIME_SECONDS", 20)
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		PostgresDSN:     os.Getenv("POSTGRES_DSN"),
		SQSQueueURL:     os.Getenv("SQS_QUEUE_URL"),
		SQSEndpoint:     os.Getenv("SQS_ENDPOINT"),
		AWSRegion:       getenv("AWS_REGION", "us-east-1"),
		BatchSize:       batchSize,
		WaitTimeSeconds: waitTimeSeconds,
	}

	if cfg.PostgresDSN == "" {
		return nil, fmt.Errorf("POSTGRES_DSN is required")
	}
	if cfg.SQSQueueURL == "" {
		return nil, fmt.Errorf("SQS_QUEUE_URL is required")
	}

	return cfg, nil
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func parseInt32(key string, fallback int32) (int32, error) {
	raw := os.Getenv(key)
	if raw == "" {
		return fallback, nil
	}

	value, err := strconv.ParseInt(raw, 10, 32)
	if err != nil || value <= 0 {
		return 0, fmt.Errorf("%s must be a positive integer", key)
	}
	return int32(value), nil
}
