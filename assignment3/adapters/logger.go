package adapters

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

type ILogger interface {
	Info(msg string, value interface{})
	Error(msg string, value interface{})
}

var _ ILogger = (*Logger)(nil)

type Logger struct {
	lg zerolog.Logger
}

func NewLogger() *Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339, NoColor: false}
	loggerCall := zerolog.New(output).With().Timestamp().Caller().Logger()
	return &Logger{loggerCall}
}

func (l *Logger) Info(msg string, value interface{}) {
	l.lg.Info().Interface(msg, value).Send()
}

func (l *Logger) Error(msg string, value interface{}) {
	l.lg.Error().Interface(msg, value).Send()
}
