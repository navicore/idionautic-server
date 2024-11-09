package main

import (
	"os"

	"github.com/navicore/idionautic-server/cmd"
	"github.com/navicore/idionautic-server/db"
	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
)

func init() {

	level, exists := os.LookupEnv("LOG_LEVEL")
	if exists {
		switch level {
		case "DEBUG":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "INFO":
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		case "WARN":
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		case "ERROR":
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		case "TRACE":
			zerolog.SetGlobalLevel(zerolog.TraceLevel)
		}
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

}

func main() {
	db.InitDB() // Initializes the SQLite DB
	if err := cmd.Execute(); err != nil {
		log.Fatal().Err(err)
	}
}
