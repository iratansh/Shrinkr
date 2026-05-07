package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "url-shortener/api/config"
    "url-shortener/api/handlers"
    "url-shortener/api/services"
)

func main() {
    cfg := config.Load()

    redisClient := services.NewRedisClient(cfg.RedisURL)
    pgClient := services.NewPostgresClient(cfg.DatabaseURL)
    sqsClient := services.NewSQSClient(cfg.SQSQueueURL, cfg.AWSRegion)

    h := handlers.New(redisClient, pgClient, sqsClient)

    r := gin.Default()
    r.POST("/shorten", h.Shorten)
    r.GET("/:code", h.Redirect)
    r.GET("/analytics/:code", h.Analytics)
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    log.Fatal(r.Run(":" + cfg.Port))
}