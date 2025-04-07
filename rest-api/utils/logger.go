package utils

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger zerolog.Logger

// InitLogger configures the global logger
func InitLogger() {
    // Pretty console output for development
    consoleWriter := zerolog.ConsoleWriter{
        Out:        os.Stdout,
        TimeFormat: time.RFC3339,
    }

    // Set global log level
    zerolog.SetGlobalLevel(zerolog.InfoLevel)
    
    // For development, you might want to use debug level
    // zerolog.SetGlobalLevel(zerolog.DebugLevel)

    // Configure logger with timestamp
    Logger = zerolog.New(consoleWriter).
        With().
        Timestamp().
        Caller().
        Logger()

    // Override the global logger
    log.Logger = Logger

    Logger.Info().Msg("Logger initialized")
}

// GetLogger returns the configured logger
func GetLogger() *zerolog.Logger {
    return &Logger
}