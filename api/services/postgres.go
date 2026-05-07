package services

import (
	"context"

	"github.com/shrinkr/api/models"
)

// PostgresClient handles persistent storage of URL records and analytics queries.
type PostgresClient struct {
	// TODO: embed *sql.DB or *pgxpool.Pool
}

// NewPostgresClient connects to Postgres using dsn.
func NewPostgresClient(dsn string) (*PostgresClient, error) {
	return &PostgresClient{}, nil
}

// InsertURL persists a new shortened URL row.
func (p *PostgresClient) InsertURL(ctx context.Context, u *models.URL) error {
	return nil
}

// GetURL fetches a URL row by short code.
func (p *PostgresClient) GetURL(ctx context.Context, code string) (*models.URL, error) {
	return nil, nil
}

// GetClickCount returns aggregated click stats for a code (used by /analytics/:code).
func (p *PostgresClient) GetClickCount(ctx context.Context, code string) (int64, error) {
	return 0, nil
}
