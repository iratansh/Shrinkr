package handlers

import "url-shortener/api/services"

// Handler owns the dependencies used by HTTP route handlers.
type Handler struct {
	redis       *services.RedisClient
	pg          *services.PostgresClient
	sqs         *services.SQSPublisher
	shortener   *services.Shortener
	baseShortURL string
}

// New wires route handlers to their backing services.
func New(
	redis *services.RedisClient,
	pg *services.PostgresClient,
	sqs *services.SQSPublisher,
	shortener *services.Shortener,
	baseShortURL string,
) *Handler {
	return &Handler{
		redis:       redis,
		pg:          pg,
		sqs:         sqs,
		shortener:   shortener,
		baseShortURL: trimTrailingSlash(baseShortURL),
	}
}

func trimTrailingSlash(value string) string {
	for len(value) > 1 && value[len(value)-1] == '/' {
		value = value[:len(value)-1]
	}
	return value
}
