package api

import (
	"fmt"
	"net/http"
)

func StartServer(iface string, port int) {
	mux := http.NewServeMux()

	// Register API routes
	mux.HandleFunc("/ingest", ingestTelemetryHandler)
	mux.HandleFunc("/analysis", getTelemetryAnalysisHandler)
	mux.HandleFunc("/telemetry", getTelemetryDataHandler)

	// Start the server
	addr := fmt.Sprintf("%s:%d", iface, port)
	log.Info().Msgf("Starting server on %s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal().Err(err).Msg("server start failed")
	}
}
