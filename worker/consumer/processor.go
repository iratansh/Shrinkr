package consumer

import (
	"context"
	"encoding/json"

	"url-shortener/worker/models"
	"url-shortener/worker/services"
)

// Processor decodes SQS messages into Click events and writes them to Postgres.
type Processor struct {
	DB *services.PostgresClient
}

// Handle decodes the raw SQS message body and persists the click.
// Return error to leave the message in the queue (will be retried / DLQ'd).
func (p *Processor) Handle(ctx context.Context, body []byte) error {
	var click models.Click
	if err := json.Unmarshal(body, &click); err != nil {
		return err
	}

	return p.DB.InsertClick(ctx, &click)
}
