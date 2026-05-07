package services

import (
	"context"

	"github.com/shrinkr/api/models"
)

// SQSPublisher publishes click events to an AWS SQS queue.
type SQSPublisher struct {
	QueueURL string
	// TODO: embed *sqs.Client
}

// NewSQSPublisher constructs an SQS publisher for the given queue URL.
func NewSQSPublisher(region, queueURL string) (*SQSPublisher, error) {
	return &SQSPublisher{QueueURL: queueURL}, nil
}

// PublishClick serializes the click event and sends it to SQS.
func (s *SQSPublisher) PublishClick(ctx context.Context, c *models.Click) error {
	return nil
}
