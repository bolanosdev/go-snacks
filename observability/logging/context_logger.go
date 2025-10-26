package logging

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

// ContextLogger wraps zerolog.Logger with trace ID context
type ContextLogger struct {
	logger  zerolog.Logger
	traceID string
}

// NewContextLogger creates a new ContextLogger with the given trace ID
// - In test mode: no output
// - In dev mode: writes to stderr with pretty formatting
// - In prod mode: writes to log file with JSON formatting
func NewContextLogger(traceID string, mode string) *ContextLogger {
	var logger zerolog.Logger

	switch mode {
	case "test":
		// No output in test mode
		logger = zerolog.Nop()
	case "dev":
		// Pretty console output for development
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).
			With().
			Timestamp().
			Caller().
			Logger()
	case "prod":
		// JSON output to file for production
		file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			// Fallback to stderr if file can't be opened
			logger = zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()
		} else {
			logger = zerolog.New(file).With().Timestamp().Caller().Logger()
		}
	default:
		// Default to stderr
		logger = zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()
	}

	return &ContextLogger{
		logger:  logger,
		traceID: traceID,
	}
}

// Event wraps zerolog.Event to automatically add trace_id
type Event struct {
	event   *zerolog.Event
	traceID string
}

// Error starts a new error-level log entry
func (cl *ContextLogger) Error() *Event {
	return &Event{
		event:   cl.logger.Error().CallerSkipFrame(1).Str("trace_id", cl.traceID),
		traceID: cl.traceID,
	}
}

// Info starts a new info-level log entry
func (cl *ContextLogger) Info() *Event {
	return &Event{
		event:   cl.logger.Info().CallerSkipFrame(1).Str("trace_id", cl.traceID),
		traceID: cl.traceID,
	}
}

// Debug starts a new debug-level log entry
func (cl *ContextLogger) Debug() *Event {
	return &Event{
		event:   cl.logger.Debug().CallerSkipFrame(1).Str("trace_id", cl.traceID),
		traceID: cl.traceID,
	}
}

// Warn starts a new warn-level log entry
func (cl *ContextLogger) Warn() *Event {
	return &Event{
		event:   cl.logger.Warn().CallerSkipFrame(1).Str("trace_id", cl.traceID),
		traceID: cl.traceID,
	}
}

// Err adds an error to the log entry
func (e *Event) Err(err error) *Event {
	e.event = e.event.Err(err)
	return e
}

// Str adds a string field to the log entry
func (e *Event) Str(key, val string) *Event {
	e.event = e.event.Str(key, val)
	return e
}

// Int adds an int field to the log entry
func (e *Event) Int(key string, val int) *Event {
	e.event = e.event.Int(key, val)
	return e
}

func (e *Event) Dur(key string, val time.Duration) *Event {
	e.event = e.event.Dur(key, val)
	return e
}

// Bool adds a bool field to the log entry
func (e *Event) Bool(key string, val bool) *Event {
	e.event = e.event.Bool(key, val)
	return e
}

// Interface adds an interface{} field to the log entry
func (e *Event) Interface(key string, val interface{}) *Event {
	e.event = e.event.Interface(key, val)
	return e
}

// WithData adds arbitrary data to the log entry
func (e *Event) WithData(key string, val interface{}) *Event {
	e.event = e.event.Interface(key, val)
	return e
}

// Msg sends the log entry with the given message
func (e *Event) Msg(msg string) {
	e.event.Msg(msg)
}

// Msgf sends the log entry with a formatted message
func (e *Event) Msgf(format string, v ...interface{}) {
	e.event.Msgf(format, v...)
}

// Send sends the log entry without a message
func (e *Event) Send() {
	e.event.Send()
}
