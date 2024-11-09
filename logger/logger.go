// logger/logger.go
package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {

	level, exists := os.LookupEnv("LOG_LEVEL")
	if exists {
		switch level {
		case "DEBUG":
			fmt.Println("DEBUG ***********")
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
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	log.Logger = logger // Set log.Logger globally within this package
}

// GetLogger returns the globally configured logger
func GetLogger() zerolog.Logger {
	return log.Logger
}
