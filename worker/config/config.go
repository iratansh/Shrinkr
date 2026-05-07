package config

// Config holds runtime configuration for the analytics worker.
type Config struct {
	PostgresDSN     string
	SQSQueueURL     string
	AWSRegion       string
	BatchSize       int32
	WaitTimeSeconds int32
}

// Load reads configuration from environment variables.
func Load() (*Config, error) {
	// TODO: read os.Getenv values, apply defaults.
	return &Config{BatchSize: 10, WaitTimeSeconds: 20}, nil
}
