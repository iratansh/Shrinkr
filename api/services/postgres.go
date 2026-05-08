package services

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"url-shortener/api/models"
)

// PostgresClient handles persistent storage of URL records and analytics queries.
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

// InsertURL persists a new shortened URL row.
func (p *PostgresClient) InsertURL(ctx context.Context, u *models.URL) error {
	_, err := p.pool.Exec(ctx, `
		INSERT INTO urls (code, long_url, expires_at)
		VALUES ($1, $2, $3)
	`, u.Code, u.LongURL, u.ExpiresAt)
	return err
}

// InsertClick appends a click event directly.
// This is used for local/dev mode when SQS is disabled.
func (p *PostgresClient) InsertClick(ctx context.Context, c *models.Click) error {
	_, err := p.pool.Exec(ctx, `
		INSERT INTO clicks (code, clicked_at, ip, user_agent, referrer, country)
		VALUES ($1, $2, NULLIF($3, '')::inet, $4, $5, $6)
	`, c.Code, c.Timestamp, c.IP, c.UserAgent, c.Referrer, c.Country)
	return err
}

// GetURL fetches a URL row by short code.
func (p *PostgresClient) GetURL(ctx context.Context, code string) (*models.URL, error) {
	var u models.URL
	err := p.pool.QueryRow(ctx, `
		SELECT code, long_url, created_at, expires_at
		FROM urls
		WHERE code = $1
		  AND (expires_at IS NULL OR expires_at > NOW())
	`, code).Scan(&u.Code, &u.LongURL, &u.CreatedAt, &u.ExpiresAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// GetClickCount returns aggregated click stats for a code (used by /analytics/:code).
func (p *PostgresClient) GetClickCount(ctx context.Context, code string) (int64, error) {
	var count int64
	err := p.pool.QueryRow(ctx, `
		SELECT COUNT(*)
		FROM clicks
		WHERE code = $1
	`, code).Scan(&count)
	return count, err
}
