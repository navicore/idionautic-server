package main

import (
	"idionautic-server/api"
	"idionautic-server/db"
	"log"
	"net/http"
)

func main() {
	db.InitDB() // Initializes the SQLite DB
	router := api.SetupRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
