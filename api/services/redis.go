package services

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisClient wraps a redis client for short-code lookups and caching.
type RedisClient struct {
	client *redis.Client
	ttl    time.Duration
}

// NewRedisClient constructs a RedisClient connected to addr.
func NewRedisClient(addr, password string) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return &RedisClient{client: client, ttl: 24 * time.Hour}, nil
}

// GetLongURL returns the long URL cached for code, or empty string on miss.
func (r *RedisClient) GetLongURL(ctx context.Context, code string) (string, error) {
	longURL, err := r.client.Get(ctx, code).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	}
	return longURL, err
}

// SetLongURL caches code -> longURL with the configured TTL.
func (r *RedisClient) SetLongURL(ctx context.Context, code, longURL string) error {
	return r.client.Set(ctx, code, longURL, r.ttl).Err()
}
