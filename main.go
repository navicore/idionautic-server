package main

import (
	"github.com/navicore/idionautic-server/db"
)

func main() {
	db.InitDB() // Initializes the SQLite DB
	// router := api.SetupRouter()
	// log.Fatal(http.ListenAndServe(":8080", router))
}
