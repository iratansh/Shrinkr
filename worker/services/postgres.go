package services

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"url-shortener/worker/models"
)

// PostgresClient handles analytics writes for the worker.
type PostgresClient struct {
	pool *pgxpool.Pool
}

// NewPostgresClient connects to Postgres using dsn.
func NewPostgresClient(dsn string) (*PostgresClient, error) {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, err
	}
	return &PostgresClient{pool: pool}, nil
}

// InsertClick appends a single click row.
func (p *PostgresClient) InsertClick(ctx context.Context, click *models.Click) error {
	_, err := p.pool.Exec(ctx, `
		INSERT INTO clicks (code, clicked_at, ip, user_agent, referrer, country)
		VALUES ($1, $2, NULLIF($3, '')::inet, $4, $5, $6)
	`, click.Code, click.Timestamp, click.IP, click.UserAgent, click.Referrer, click.Country)
	return err
}
