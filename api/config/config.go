package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds runtime configuration loaded from environment variables.
type Config struct {
	HTTPPort       string
	RedisAddr      string
	RedisPassword  string
	PostgresDSN    string
	SQSQueueURL    string
	SQSEndpoint    string
	AWSRegion      string
	BaseShortURL   string
	CORSOrigins    string
	ShortCodeBytes int
}

// Load reads configuration from environment variables and returns a Config.
func Load() (*Config, error) {
	codeBytes, err := strconv.Atoi(getenv("SHORT_CODE_BYTES", "6"))
	if err != nil || codeBytes <= 0 {
		return nil, fmt.Errorf("SHORT_CODE_BYTES must be a positive integer")
	}

	cfg := &Config{
		HTTPPort:       getenv("HTTP_PORT", "8080"),
		RedisAddr:      getenv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:  os.Getenv("REDIS_PASSWORD"),
		PostgresDSN:    os.Getenv("POSTGRES_DSN"),
		SQSQueueURL:    os.Getenv("SQS_QUEUE_URL"),
		SQSEndpoint:    os.Getenv("SQS_ENDPOINT"),
		AWSRegion:      getenv("AWS_REGION", "us-east-1"),
		BaseShortURL:   getenv("BASE_SHORT_URL", "http://localhost:8080"),
		CORSOrigins:    getenv("CORS_ALLOWED_ORIGINS", "http://localhost:5173"),
		ShortCodeBytes: codeBytes,
	}

	if cfg.PostgresDSN == "" {
		return nil, fmt.Errorf("POSTGRES_DSN is required")
	}

	return cfg, nil
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
