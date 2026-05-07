package models

import "time"

// Click is the message payload published by the API for each redirect.
type Click struct {
	Code      string    `json:"code"`
	Timestamp time.Time `json:"timestamp"`
	IP        string    `json:"ip,omitempty"`
	UserAgent string    `json:"user_agent,omitempty"`
	Referrer  string    `json:"referrer,omitempty"`
	Country   string    `json:"country,omitempty"`
}
