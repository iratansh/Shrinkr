package consumer

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"url-shortener/worker/services"
)

type SQSConsumer struct {
	client          *sqs.Client
	queueURL        string
	batchSize       int32
	waitTimeSeconds int32
	processor       *Processor
}

func NewSQSConsumer(queueURL, region string, batchSize, waitTimeSeconds int32, pg *services.PostgresClient) (*SQSConsumer, error) {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	if err != nil {
		return nil, err
	}

	return &SQSConsumer{
		client:          sqs.NewFromConfig(cfg),
		queueURL:        queueURL,
		batchSize:       batchSize,
		waitTimeSeconds: waitTimeSeconds,
		processor:       &Processor{DB: pg},
	}, nil
}

func (c *SQSConsumer) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		result, err := c.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
			QueueUrl:            &c.queueURL,
			MaxNumberOfMessages: c.batchSize,
			WaitTimeSeconds:     c.waitTimeSeconds,
		})
		if err != nil {
			log.Printf("receive messages: %v", err)
			continue
		}

		for _, msg := range result.Messages {
			if msg.Body == nil || msg.ReceiptHandle == nil {
				continue
			}

			if err := c.processor.Handle(ctx, []byte(*msg.Body)); err != nil {
				log.Printf("process click event: %v", err)
				continue
			}

			if _, err := c.client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
				QueueUrl:      &c.queueURL,
				ReceiptHandle: msg.ReceiptHandle,
			}); err != nil {
				log.Printf("delete processed message: %v", err)
			}
		}
	}
}