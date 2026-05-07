package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"url-shortener/api/config"
	"url-shortener/api/handlers"
	"url-shortener/api/services"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	redisClient, err := services.NewRedisClient(cfg.RedisAddr, cfg.RedisPassword)
	if err != nil {
		log.Fatalf("connect redis: %v", err)
	}

	pgClient, err := services.NewPostgresClient(cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("connect postgres: %v", err)
	}

	sqsClient, err := services.NewSQSPublisher(cfg.AWSRegion, cfg.SQSQueueURL)
	if err != nil {
		log.Fatalf("connect sqs: %v", err)
	}

	h := handlers.New(
		redisClient,
		pgClient,
		sqsClient,
		services.NewShortener(cfg.ShortCodeBytes),
		cfg.BaseShortURL,
	)

	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.POST("/shorten", h.Shorten)
	r.GET("/analytics/:code", h.Analytics)
	r.GET("/:code", h.Redirect)

	log.Fatal(r.Run(":" + cfg.HTTPPort))
}