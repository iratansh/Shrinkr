package models

import "time"

// URL represents a shortened URL record stored in Postgres and cached in Redis.
type URL struct {
	Code      string    `json:"code"        db:"code"`
	LongURL   string    `json:"long_url"    db:"long_url"`
	CreatedAt time.Time `json:"created_at"  db:"created_at"`
	ExpiresAt time.Time `json:"expires_at"  db:"expires_at"`
}
