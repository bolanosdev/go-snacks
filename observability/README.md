# Observability

Observability utilities for logging and tracing.

## Installation

```bash
go get github.com/bolanosdev/go-snacks/observability
```

## Logging

### ContextLogger

A wrapper around zerolog that enforces adding a `trace_id` string to all log entries.

#### Usage

```go
import "github.com/bolanosdev/go-snacks/observability/logging"

// Create logger with trace ID
logger := logging.NewContextLogger("trace-123", "dev")

// Log messages - trace_id is automatically included
logger.Info().Msg("operation started")
logger.Error().Err(err).Msg("operation failed")
logger.Debug().Str("user", "john").Int("count", 42).Msg("user action")
logger.Warn().Msgf("threshold exceeded: %d", value)

// All log entries automatically include trace_id field
```

#### Modes

- `test`: No output (useful for testing)
- `dev`: Pretty console output to stderr with timestamp and caller info
- `prod`: JSON output to `app.log` file with timestamp and caller info
- Any other value: Defaults to JSON output to stderr

#### Available Methods

**Log Levels:**
- `Info()` - Info-level log entry
- `Error()` - Error-level log entry
- `Debug()` - Debug-level log entry
- `Warn()` - Warn-level log entry

**Field Methods:**
- `Err(error)` - Add error to log entry
- `Str(key, val)` - Add string field
- `Int(key, val)` - Add int field
- `Dur(key, val)` - Add duration field
- `Bool(key, val)` - Add bool field
- `Interface(key, val)` - Add interface{} field
- `WithData(key, val)` - Add arbitrary data (alias for Interface)

**Output Methods:**
- `Msg(msg)` - Send log with message
- `Msgf(format, ...args)` - Send log with formatted message
- `Send()` - Send log without message

## Running Tests

```bash
go test ./observability/...
```
