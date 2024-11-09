package models

import "time"

type TelemetryData struct {
	EventType string    `json:"eventType"`
	Target    string    `json:"target"`
	Count     int       `json:"count"`
	Timestamp time.Time `json:"timestamp"` // Use time.Time directly
}
