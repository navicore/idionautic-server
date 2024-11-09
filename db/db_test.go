package db

import (
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/navicore/idionautic-server/models"
)

// setupTestDB initializes a temporary SQLite database for testing
func setupTestDB(t *testing.T) {
	var err error
	DB, err = sql.Open("sqlite3", "./test_telemetry.db")
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}

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
		t.Fatalf("Failed to create test table: %v", err)
	}
}

// teardownTestDB removes the temporary test database
func teardownTestDB() {
	DB.Close()
	os.Remove("./test_telemetry.db")
}

// TestSaveTelemetryData verifies that telemetry data is saved to the database correctly
func TestSaveTelemetryData(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB()
	currentTime := time.Now()
	testData := models.TelemetryData{
		EventType: "functionCall",
		Target:    "calculateMetrics",
		Count:     10,
		Timestamp: currentTime,
	}

	if err := SaveTelemetryData(testData); err != nil {
		t.Fatalf("Failed to save telemetry data: %v", err)
	}

	// Query the database to verify the data was saved
	var id int
	var eventType, target string
	var count int
	var timestamp_int int

	err := DB.QueryRow(`
        SELECT id, event_type, target, count, timestamp
        FROM telemetry
        WHERE event_type = ? AND target = ?
    `, testData.EventType, testData.Target).Scan(&id, &eventType, &target, &count, &timestamp_int)

	if err != nil {
		t.Fatalf("Failed to retrieve telemetry data: %v", err)
	}

	timestamp := time.UnixMilli(int64(timestamp_int)).Truncate(time.Millisecond)
	testDataTimestamp := testData.Timestamp.Truncate(time.Millisecond)

	if eventType != testData.EventType || target != testData.Target || count != testData.Count || timestamp != testDataTimestamp {
		t.Errorf("Retrieved data does not match: got (%s, %s, %d, %s), want (%s, %s, %d, %s)",
			eventType, target, count, timestamp,
			testData.EventType, testData.Target, testData.Count, testData.Timestamp)
	}
}
