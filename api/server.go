package api

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer(iface string, port int) {
	mux := http.NewServeMux()

	// Register API routes
	mux.HandleFunc("/ingest", ingestTelemetryHandler)

	// Start the server
	addr := fmt.Sprintf("%s:%d", iface, port)
	log.Printf("Starting server on %s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
