package services

import "context"

// RedisClient wraps a redis client for short-code lookups and caching.
type RedisClient struct {
	// TODO: embed *redis.Client
}

// NewRedisClient constructs a RedisClient connected to addr.
func NewRedisClient(addr, password string) (*RedisClient, error) {
	return &RedisClient{}, nil
}

// GetLongURL returns the long URL cached for code, or empty string on miss.
func (r *RedisClient) GetLongURL(ctx context.Context, code string) (string, error) {
	return "", nil
}

// SetLongURL caches code -> longURL with the configured TTL.
func (r *RedisClient) SetLongURL(ctx context.Context, code, longURL string) error {
	return nil
}
