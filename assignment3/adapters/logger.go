package adapters

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

type Logger interface {
	Info(msg string, value interface{})
	Error(msg string, value interface{})
}

type logger struct {
	zerolog.Logger
}

func NewLogger() Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339, NoColor: false}
	loggerCall := zerolog.New(output).With().Timestamp().Caller().Logger()
	return &logger{loggerCall}
}

func (l *logger) Info(msg string, value interface{}) {
	l.Logger.Info().Interface(msg, value).Send()
}

func (l *logger) Error(msg string, value interface{}) {
	l.Logger.Error().Interface(msg, value).Send()
}
