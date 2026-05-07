package main

import (
    "log"
    "url-shortener/worker/config"
    "url-shortener/worker/consumer"
    "url-shortener/worker/services"
)

func main() {
    cfg := config.Load()

    pgClient := services.NewPostgresClient(cfg.DatabaseURL)
    sqsConsumer := consumer.NewSQSConsumer(
        cfg.SQSQueueURL,
        cfg.AWSRegion,
        cfg.BatchSize,
        pgClient,
    )

    log.Println("Worker started, polling SQS...")
    sqsConsumer.Start()
}