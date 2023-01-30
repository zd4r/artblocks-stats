package logger

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

func New(level string, structured bool) *zerolog.Logger {
	switch strings.ToLower(level) {
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	var output io.Writer
	output = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	if structured {
		output = os.Stdout
	}

	logger := zerolog.New(output).With().Timestamp().Logger()

	return &logger
}
