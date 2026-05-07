package services

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"url-shortener/api/models"
)

// SQSPublisher publishes click events to an AWS SQS queue.
type SQSPublisher struct {
	QueueURL string
	client   *sqs.Client
}

// NewSQSPublisher constructs an SQS publisher for the given queue URL.
func NewSQSPublisher(region, queueURL string) (*SQSPublisher, error) {
	if queueURL == "" {
		return &SQSPublisher{}, nil
	}

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	if err != nil {
		return nil, err
	}

	return &SQSPublisher{
		QueueURL: queueURL,
		client:   sqs.NewFromConfig(cfg),
	}, nil
}

// PublishClick serializes the click event and sends it to SQS.
func (s *SQSPublisher) PublishClick(ctx context.Context, c *models.Click) error {
	if s.QueueURL == "" || s.client == nil {
		return nil
	}

	body, err := json.Marshal(c)
	if err != nil {
		return err
	}

	_, err = s.client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    &s.QueueURL,
		MessageBody: awsString(string(body)),
	})
	return err
}

func awsString(value string) *string {
	return &value
}
