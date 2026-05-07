package config

// Config holds runtime configuration loaded from environment variables.
type Config struct {
	HTTPPort       string
	RedisAddr      string
	RedisPassword  string
	PostgresDSN    string
	SQSQueueURL    string
	AWSRegion      string
	BaseShortURL   string
	ShortCodeBytes int
}

// Load reads configuration from environment variables and returns a Config.
func Load() (*Config, error) {
	// TODO: read os.Getenv values, apply defaults, validate required fields.
	return &Config{}, nil
}
