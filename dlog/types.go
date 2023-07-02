package dlog

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	Err   zerolog.Logger
	Trace zerolog.Logger
}

func NewLogger() *Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	return &Logger{Err: zerolog.New(os.Stderr).With().Str("service", "dolk").
		Timestamp().
		Logger(), Trace: zerolog.New(os.Stdout).With().Str("service", "dolk").
		Timestamp().
		Logger()}
}
