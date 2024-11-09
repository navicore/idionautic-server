package main

import (
	"github.com/navicore/idionautic-server/cmd"
	"github.com/navicore/idionautic-server/db"
	"github.com/navicore/idionautic-server/logger"
	"github.com/rs/zerolog"
)

var log zerolog.Logger

func init() {
	log = logger.GetLogger()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func main() {
	log.Debug().Msg("Logger initialized with level " + zerolog.GlobalLevel().String())
	db.InitDB() // Initializes the SQLite DB
	if err := cmd.Execute(); err != nil {
		log.Fatal().Err(err)
	}
}
