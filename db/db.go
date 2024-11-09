// db/db.go

package db

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/navicore/idionautic-server/logger"
	"github.com/navicore/idionautic-server/models"
)

var log = logger.GetLogger()
var DB *sql.DB

// InitDB initializes the SQLite database
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./telemetry.db")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open database")
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
		log.Fatal().Err(err).Msg("Failed to create table")
	}
}

// SaveTelemetryData saves a telemetry record in the database
func SaveTelemetryData(data models.TelemetryData) error {
	// Convert the time.Time to Unix milliseconds
	timestampMillis := data.Timestamp.UnixMilli()

	_, err := DB.Exec(`
        INSERT INTO telemetry (event_type, target, count, timestamp)
        VALUES (?, ?, ?, ?)
    `, data.EventType, data.Target, data.Count, timestampMillis)
	log.Debug().Msg("Saved telemetry data with timestamp in milliseconds")
	return err
}

func GetTelemetryAnalysis() (models.AnalysisResponse, error) {
	var analysis models.AnalysisResponse
	var startTimestamp, endTimestamp int64

	// Get the start and end times as Unix timestamps
	err := DB.QueryRow("SELECT MIN(timestamp), MAX(timestamp) FROM telemetry").Scan(&startTimestamp, &endTimestamp)
	if err != nil {
		log.Err(err).Msg("Error retrieving time period")
		return analysis, err
	}

	// Convert to ISO 8601 format
	analysis.StartTime = time.Unix(0, startTimestamp*int64(time.Millisecond)).UTC().Format(time.RFC3339)
	analysis.EndTime = time.Unix(0, endTimestamp*int64(time.Millisecond)).UTC().Format(time.RFC3339)

	// Query counts of each eventType
	rows, err := DB.Query("SELECT event_type, COUNT(*) FROM telemetry GROUP BY event_type")
	if err != nil {
		log.Err(err).Msg("Error retrieving eventType counts")
		return analysis, err
	}
	defer rows.Close()

	analysis.EventTypeStats = make(map[string]int)
	for rows.Next() {
		var eventType string
		var count int
		if err := rows.Scan(&eventType, &count); err != nil {
			log.Err(err).Msg("Error scanning eventType row")
			return analysis, err
		}
		analysis.EventTypeStats[eventType] = count
	}

	// Query counts of each target
	rows, err = DB.Query("SELECT target, COUNT(*) FROM telemetry GROUP BY target")
	if err != nil {
		log.Err(err).Msg("Error retrieving target counts")
		return analysis, err
	}
	defer rows.Close()

	analysis.TargetStats = make(map[string]int)
	for rows.Next() {
		var target string
		var count int
		if err := rows.Scan(&target, &count); err != nil {
			log.Err(err).Msg("Error scanning target row")
			return analysis, err
		}
		analysis.TargetStats[target] = count
	}

	return analysis, nil
}

func GetPaginatedTelemetryData(limit, offset int) ([]models.TelemetryData, error) {
	rows, err := DB.Query(`
        SELECT event_type, target, count, timestamp
        FROM telemetry
        ORDER BY id
        LIMIT ? OFFSET ?
    `, limit, offset)

	if err != nil {
		log.Printf("Error querying telemetry data: %v", err)
		return nil, err
	}
	defer rows.Close()

	var telemetryData []models.TelemetryData
	for rows.Next() {
		var data models.TelemetryData
		var timestamp int64

		if err := rows.Scan(&data.EventType, &data.Target, &data.Count, &timestamp); err != nil {
			log.Printf("Error scanning telemetry data row: %v", err)
			return nil, err
		}

		// Convert the integer timestamp to time.Time and store as ISO 8601
		data.Timestamp = time.UnixMilli(timestamp).UTC()

		telemetryData = append(telemetryData, data)
	}

	return telemetryData, nil
}
