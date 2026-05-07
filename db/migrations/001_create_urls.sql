CREATE TABLE IF NOT EXISTS urls (
    code        VARCHAR(16) PRIMARY KEY,
    long_url    TEXT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at  TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_urls_created_at ON urls(created_at DESC);