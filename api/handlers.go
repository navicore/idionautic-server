// api/handlers.go

package api

import (
	"encoding/json"
	"net/http"

	"github.com/navicore/idionautic-server/db"
)

func ingestTelemetryHandler(w http.ResponseWriter, r *http.Request) {
	var data db.TelemetryData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Store telemetry data in SQLite
	if err := db.SaveTelemetryData(data); err != nil {
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
