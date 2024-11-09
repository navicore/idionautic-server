// api/handlers.go

package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/navicore/idionautic-server/db"
	"github.com/navicore/idionautic-server/models"
)

func ingestTelemetryHandler(w http.ResponseWriter, r *http.Request) {
	var data models.TelemetryData

	// Log raw incoming JSON for clarity
	rawData, _ := io.ReadAll(r.Body)

	// Parse JSON payload into TelemetryData struct
	if err := json.Unmarshal(rawData, &data); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Pass parsed data to database saving function
	if err := db.SaveTelemetryData(data); err != nil {
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func getTelemetryAnalysisHandler(w http.ResponseWriter, _ *http.Request) {
	analysis, err := db.GetTelemetryAnalysis()
	if err != nil {
		http.Error(w, "Failed to retrieve analysis data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(analysis)
	if err != nil {
		http.Error(w, "Failed to retrieve analysis data", http.StatusInternalServerError)
		return
	}
}

func getTelemetryDataHandler(w http.ResponseWriter, r *http.Request) {
	// Default limit and offset values
	limit := 10
	offset := 0

	// Parse the "limit" query parameter
	if l := r.URL.Query().Get("limit"); l != "" {
		if parsedLimit, err := strconv.Atoi(l); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	// Parse the "offset" query parameter
	if o := r.URL.Query().Get("offset"); o != "" {
		if parsedOffset, err := strconv.Atoi(o); err == nil && parsedOffset >= 0 {
			offset = parsedOffset
		}
	}

	// Fetch paginated telemetry data from the database
	telemetryData, err := db.GetPaginatedTelemetryData(limit, offset)
	if err != nil {
		http.Error(w, "Failed to retrieve telemetry data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(telemetryData)
	if err != nil {
		http.Error(w, "Failed to retrieve telemetry data", http.StatusInternalServerError)
		return
	}
}
