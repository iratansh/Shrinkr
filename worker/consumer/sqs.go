package consumer

import (
    "context"
    "encoding/json"
    "log"
    "github.com/aws/aws-sdk-go-v2/service/sqs"
    "url-shortener/worker/models"
    "url-shortener/worker/services"
)

type SQSConsumer struct {
    client    *sqs.Client
    queueURL  string
    batchSize int32
    pg        *services.PostgresClient
}

func (c *SQSConsumer) Start() {
    for {
        result, err := c.client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
            QueueUrl:            &c.queueURL,
            MaxNumberOfMessages: c.batchSize,
            WaitTimeSeconds:     20, // Long polling
        })
        if err != nil {
            log.Printf("Error receiving messages: %v", err)
            continue
        }

        for _, msg := range result.Messages {
            var event models.ClickEvent
            if err := json.Unmarshal([]byte(*msg.Body), &event); err != nil {
                log.Printf("Error unmarshaling event: %v", err)
                continue
            }

            if err := c.pg.SaveClick(&event); err != nil {
                log.Printf("Error saving click: %v", err)
                continue
            }

            // Delete message from queue after successful processing
            c.client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
                QueueUrl:      &c.queueURL,
                ReceiptHandle: msg.ReceiptHandle,
            })
        }
    }
}