// db/db.go

package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// TelemetryData represents a telemetry record
type TelemetryData struct {
	EventType string
	Target    string
	Count     int
	Timestamp int64
}

// InitDB initializes the SQLite database
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./telemetry.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create table if it doesn't exist
	_, err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS telemetry (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            event_type TEXT,
            target TEXT,
            count INTEGER,
            timestamp INTEGER
        )
    `)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

// SaveTelemetryData saves a telemetry record in the database
func SaveTelemetryData(data TelemetryData) error {
	_, err := DB.Exec(`
        INSERT INTO telemetry (event_type, target, count, timestamp)
        VALUES (?, ?, ?, ?)
    `, data.EventType, data.Target, data.Count, data.Timestamp)
	return err
}
