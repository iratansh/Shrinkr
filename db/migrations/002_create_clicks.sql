CREATE TABLE clicks (
    id          SERIAL PRIMARY KEY,
    short_code  VARCHAR(10) NOT NULL,
    user_agent  TEXT,
    ip_address  VARCHAR(45),
    referrer    TEXT,
    clicked_at  TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_clicks_short_code ON clicks(short_code);