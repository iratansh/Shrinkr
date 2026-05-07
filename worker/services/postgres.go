package services

import "context"

// PostgresClient handles analytics writes for the worker.
type PostgresClient struct {
	// TODO: embed *sql.DB or *pgxpool.Pool
}

// NewPostgresClient connects to Postgres using dsn.
func NewPostgresClient(dsn string) (*PostgresClient, error) {
	return &PostgresClient{}, nil
}

// InsertClick appends a single click row.
// TODO: accept *models.Click once a shared models package or vendoring strategy is decided.
func (p *PostgresClient) InsertClick(ctx context.Context, payload []byte) error {
	return nil
}
