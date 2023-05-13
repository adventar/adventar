package util

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Logger = zerolog.New(
	zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.DateTime,
	},
).With().Timestamp().Logger()

func SetLogLevel(level string) {
	l, err := zerolog.ParseLevel(level)
	if err != nil {
		Logger.Fatal().Err(err).Msg("")
	}
	Logger = Logger.Level(l)
}
