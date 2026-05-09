package main

import (
	"log"
	"net/http"
	"strings"

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
	r.Use(corsMiddleware(cfg.CORSOrigins))
	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.POST("/shorten", h.Shorten)
	r.GET("/analytics/:code", h.Analytics)
	r.GET("/:code", h.Redirect)

	log.Fatal(r.Run(":" + cfg.HTTPPort))
}

func corsMiddleware(allowedOrigins string) gin.HandlerFunc {
	allowed := map[string]struct{}{}
	for _, origin := range strings.Split(allowedOrigins, ",") {
		origin = strings.TrimSpace(origin)
		if origin != "" {
			allowed[origin] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if _, ok := allowed[origin]; ok {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Vary", "Origin")
			c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
