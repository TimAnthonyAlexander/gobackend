package models

import "time"

// DateTimeResponse represents the date and time information returned by the API
type DateTimeResponse struct {
	CurrentTime string `json:"current_time"`
	Timestamp   int64  `json:"timestamp"`
	UTC         string `json:"utc"`
}

// NewDateTimeResponse creates a new DateTimeResponse with the current date and time
func NewDateTimeResponse() DateTimeResponse {
	now := time.Now()
	return DateTimeResponse{
		CurrentTime: now.Format(time.RFC3339),
		Timestamp:   now.Unix(),
		UTC:         now.UTC().Format(time.RFC3339),
	}
} 