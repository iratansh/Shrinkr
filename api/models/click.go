package models

import "time"

// Click represents a single redirect event published to SQS and persisted by the worker.
type Click struct {
	Code      string    `json:"code"`
	Timestamp time.Time `json:"timestamp"`
	IP        string    `json:"ip,omitempty"`
	UserAgent string    `json:"user_agent,omitempty"`
	Referrer  string    `json:"referrer,omitempty"`
	Country   string    `json:"country,omitempty"`
}
