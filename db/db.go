package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./telemetry.db")
	if err != nil {
		log.Fatal(err)
	}
	// Execute schema.sql to initialize tables
}
