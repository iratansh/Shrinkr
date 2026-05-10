package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"url-shortener/worker/config"
	"url-shortener/worker/consumer"
	"url-shortener/worker/services"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	pgClient, err := services.NewPostgresClient(cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("connect postgres: %v", err)
	}

	sqsConsumer, err := consumer.NewSQSConsumer(
		cfg.SQSQueueURL,
		cfg.AWSRegion,
		cfg.SQSEndpoint,
		cfg.BatchSize,
		cfg.WaitTimeSeconds,
		pgClient,
	)
	if err != nil {
		log.Fatalf("connect sqs: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.Println("worker started, polling SQS...")
	sqsConsumer.Start(ctx)
}
