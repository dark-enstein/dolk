package dlog

import (
	"bytes"
	"io"
	"os"

	"github.com/rs/zerolog"
)

// Buffer is a dlog native buffer container that could contain series of byte
// information that could be flushed to a logging backend or filesystem when
// the duration of a request is complete.
// See TraceBuffer and Bufferer interface for more details on its implementation
type Buffer bytes.Buffer

// TraceBuffer is a buffer container that could contain series of log/tracing
// information that would be flushed to a logging backend or filesystem when
// the duration of a request is complete.
type TraceBuffer Buffer

// NewTraceBuffer sets up a TraceBuffer instance,
// however this function returns a struct literal who hasn't been assigned a
// memory address yet. To assign memory address and begin to use the
// TraceBuffer instance see the TraceBuffer.Activate method.
//
// For now the activate logic would be wrapped by the caller implementation,
// however, this requires some more thought. TODO
func NewTraceBuffer() *TraceBuffer {
	t := &TraceBuffer{}
	return t
}

// Activate method assigns the TraceBuffer a memory address by converting it
// into a bytes.Buffer using type conversion. This allows it to access bytes.Buffer
// own methods, alongside its Bufferer methods.
func (t *TraceBuffer) Activate() *bytes.Buffer {
	b := bytes.Buffer(*t)
	return &b
}

// The Bufferer interface defines the set of functions needed to process the
// TraceBuffer.
type Bufferer interface { // TODO
	MarshToJsonByte() []byte
	MarshToString() string
	MarshToFS(location string) bool
	Activate() *bytes.Buffer
}

// A Logger represents the main dlog object that generates lines of json to
// output to the io.Writer.
type Logger struct {
	sweep bool // flag for activating tracing for server
	pot   *bytes.Buffer
}

// NewLogger generates a Logger with the default configuration for dolk.
// If sweep isn't set, tracing isn't activated on the server
func NewLogger(sweep bool) *Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	return &Logger{
		sweep: sweep,
		pot:   NewTraceBuffer().Activate(),
	}
}

func (l *Logger) Err() *zerolog.Logger {
	return loggingObject(os.Stderr)
}

func (l *Logger) Trace() *zerolog.Logger {
	return loggingObject(os.Stdout)
}

// loggingObject returns a zerolog event configured to write logs to the
// specified io.Writer interface.
func loggingObject(writer io.Writer) *zerolog.Logger {
	z := zerolog.New(writer).With().Str("service", "dolk").
		Timestamp().
		Logger()
	return &z
}

//
//func (l *Logger) Add(event *zerolog.Event) (int, error) {
//	l.pot.Write(l.)
//}
