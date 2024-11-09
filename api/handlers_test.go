package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/navicore/idionautic-server/db"
	"github.com/navicore/idionautic-server/models"
)

// Initialize the database for testing
func init() {
	db.InitDB() // Initialize SQLite for testing purposes
}

// Helper function to create a JSON request
func createRequest(t *testing.T, method, url string, body interface{}) *http.Request {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	return req
}

// Test successful ingestion of telemetry data
func TestIngestTelemetryHandler_Success(t *testing.T) {
	currentTime := time.Now()
	// Sample telemetry data
	data := models.TelemetryData{
		EventType: "functionCall",
		Target:    "calculateMetrics",
		Count:     10,
		Timestamp: currentTime,
	}

	// Set up the request and response recorder
	req := createRequest(t, "POST", "/ingest", data)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ingestTelemetryHandler)

	// Invoke the handler
	handler.ServeHTTP(rr, req)

	// Check the response code
	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusAccepted)
	}
}

// Test for invalid JSON payload
func TestIngestTelemetryHandler_InvalidJSON(t *testing.T) {
	req, err := http.NewRequest("POST", "/ingest", bytes.NewBuffer([]byte("invalid json")))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ingestTelemetryHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}
