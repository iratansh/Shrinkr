package consumer

import (
	"context"

	"github.com/shrinkr/worker/services"
)

// Processor decodes SQS messages into Click events and writes them to Postgres.
type Processor struct {
	DB *services.PostgresClient
}

// Handle decodes the raw SQS message body and persists the click.
// Return error to leave the message in the queue (will be retried / DLQ'd).
func (p *Processor) Handle(ctx context.Context, body []byte) error {
	// TODO: json.Unmarshal into models.Click, then DB.InsertClick.
	return nil
}
