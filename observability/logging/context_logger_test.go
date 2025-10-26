package logging

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
)

func TestNewContextLogger(t *testing.T) {
	tests := []struct {
		name    string
		traceID string
		mode    string
	}{
		{
			name:    "test mode",
			traceID: "test-trace-123",
			mode:    "test",
		},
		{
			name:    "dev mode",
			traceID: "dev-trace-456",
			mode:    "dev",
		},
		{
			name:    "default mode",
			traceID: "default-trace-789",
			mode:    "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := NewContextLogger(tt.traceID, tt.mode)
			if logger == nil {
				t.Fatal("expected logger to be non-nil")
			}
			if logger.traceID != tt.traceID {
				t.Errorf("expected traceID %s, got %s", tt.traceID, logger.traceID)
			}
		})
	}
}

func TestNewContextLoggerProdMode(t *testing.T) {
	traceID := "prod-trace-123"
	logger := NewContextLogger(traceID, "prod")

	if logger == nil {
		t.Fatal("expected logger to be non-nil")
	}
	if logger.traceID != traceID {
		t.Errorf("expected traceID %s, got %s", traceID, logger.traceID)
	}

	if _, err := os.Stat("app.log"); err == nil {
		os.Remove("app.log")
	}
}

func TestContextLoggerWithBuffer(t *testing.T) {
	var buf bytes.Buffer
	traceID := "test-trace-456"

	logger := &ContextLogger{
		logger:  zerolog.New(&buf).With().Timestamp().Logger(),
		traceID: traceID,
	}

	t.Run("Info with trace_id", func(t *testing.T) {
		buf.Reset()
		logger.Info().Msg("test info message")

		var logEntry map[string]interface{}
		if err := json.Unmarshal(buf.Bytes(), &logEntry); err != nil {
			t.Fatalf("failed to parse log output: %v", err)
		}

		if logEntry["trace_id"] != traceID {
			t.Errorf("expected trace_id %s, got %v", traceID, logEntry["trace_id"])
		}
		if logEntry["level"] != "info" {
			t.Errorf("expected level info, got %v", logEntry["level"])
		}
		if logEntry["message"] != "test info message" {
			t.Errorf("expected message 'test info message', got %v", logEntry["message"])
		}
	})

	t.Run("Error with trace_id", func(t *testing.T) {
		buf.Reset()
		logger.Error().Msg("test error message")

		var logEntry map[string]interface{}
		if err := json.Unmarshal(buf.Bytes(), &logEntry); err != nil {
			t.Fatalf("failed to parse log output: %v", err)
		}

		if logEntry["trace_id"] != traceID {
			t.Errorf("expected trace_id %s, got %v", traceID, logEntry["trace_id"])
		}
		if logEntry["level"] != "error" {
			t.Errorf("expected level error, got %v", logEntry["level"])
		}
	})

	t.Run("Debug with trace_id", func(t *testing.T) {
		buf.Reset()
		logger.Debug().Msg("test debug message")

		var logEntry map[string]interface{}
		if err := json.Unmarshal(buf.Bytes(), &logEntry); err != nil {
			t.Fatalf("failed to parse log output: %v", err)
		}

		if logEntry["trace_id"] != traceID {
			t.Errorf("expected trace_id %s, got %v", traceID, logEntry["trace_id"])
		}
		if logEntry["level"] != "debug" {
			t.Errorf("expected level debug, got %v", logEntry["level"])
		}
	})

	t.Run("Warn with trace_id", func(t *testing.T) {
		buf.Reset()
		logger.Warn().Msg("test warn message")

		var logEntry map[string]interface{}
		if err := json.Unmarshal(buf.Bytes(), &logEntry); err != nil {
			t.Fatalf("failed to parse log output: %v", err)
		}

		if logEntry["trace_id"] != traceID {
			t.Errorf("expected trace_id %s, got %v", traceID, logEntry["trace_id"])
		}
		if logEntry["level"] != "warn" {
			t.Errorf("expected level warn, got %v", logEntry["level"])
		}
	})
}

func TestEventMethods(t *testing.T) {
	var buf bytes.Buffer
	traceID := "test-trace-789"

	logger := &ContextLogger{
		logger:  zerolog.New(&buf).With().Timestamp().Logger(),
		traceID: traceID,
	}

	t.Run("Err method", func(t *testing.T) {
		buf.Reset()
		testErr := errors.New("test error")
		logger.Error().Err(testErr).Msg("error occurred")

		var logEntry map[string]interface{}
		if err := json.Unmarshal(buf.Bytes(), &logEntry); err != nil {
			t.Fatalf("failed to parse log output: %v", err)
		}

		if logEntry["error"] != testErr.Error() {
			t.Errorf("expected error %s, got %v", testErr.Error(), logEntry["error"])
		}
	})

	t.Run("Str method", func(t *testing.T) {
		buf.Reset()
		logger.Info().Str("user", "john").Msg("user action")

		var logEntry map[string]interface{}
		if err := json.Unmarshal(buf.Bytes(), &logEntry); err != nil {
			t.Fatalf("failed to parse log output: %v", err)
		}

		if logEntry["user"] != "john" {
			t.Errorf("expected user john, got %v", logEntry["user"])
		}
	})

	t.Run("Int method", func(t *testing.T) {
		buf.Reset()
		logger.Info().Int("count", 42).Msg("counted items")

		var logEntry map[string]interface{}
		if err := json.Unmarshal(buf.Bytes(), &logEntry); err != nil {
			t.Fatalf("failed to parse log output: %v", err)
		}

		if logEntry["count"].(float64) != 42 {
			t.Errorf("expected count 42, got %v", logEntry["count"])
		}
	})

	t.Run("Dur method", func(t *testing.T) {
		buf.Reset()
		duration := 5 * time.Second
		logger.Info().Dur("duration", duration).Msg("operation completed")

		var logEntry map[string]interface{}
		if err := json.Unmarshal(buf.Bytes(), &logEntry); err != nil {
			t.Fatalf("failed to parse log output: %v", err)
		}

		if logEntry["duration"].(float64) != 5000 {
			t.Errorf("expected duration 5000, got %v", logEntry["duration"])
		}
	})

	t.Run("Bool method", func(t *testing.T) {
		buf.Reset()
		logger.Info().Bool("success", true).Msg("operation status")

		var logEntry map[string]interface{}
		if err := json.Unmarshal(buf.Bytes(), &logEntry); err != nil {
			t.Fatalf("failed to parse log output: %v", err)
		}

		if logEntry["success"] != true {
			t.Errorf("expected success true, got %v", logEntry["success"])
		}
	})

	t.Run("Interface method", func(t *testing.T) {
		buf.Reset()
		data := map[string]string{"key": "value"}
		logger.Info().Interface("data", data).Msg("with data")

		var logEntry map[string]interface{}
		if err := json.Unmarshal(buf.Bytes(), &logEntry); err != nil {
			t.Fatalf("failed to parse log output: %v", err)
		}

		dataMap := logEntry["data"].(map[string]interface{})
		if dataMap["key"] != "value" {
			t.Errorf("expected data.key to be 'value', got %v", dataMap["key"])
		}
	})

	t.Run("WithData method", func(t *testing.T) {
		buf.Reset()
		data := map[string]int{"count": 10}
		logger.Info().WithData("metrics", data).Msg("metrics logged")

		var logEntry map[string]interface{}
		if err := json.Unmarshal(buf.Bytes(), &logEntry); err != nil {
			t.Fatalf("failed to parse log output: %v", err)
		}

		metricsMap := logEntry["metrics"].(map[string]interface{})
		if metricsMap["count"].(float64) != 10 {
			t.Errorf("expected metrics.count to be 10, got %v", metricsMap["count"])
		}
	})

	t.Run("Msgf method", func(t *testing.T) {
		buf.Reset()
		logger.Info().Msgf("user %s logged in at %d", "alice", 12345)

		var logEntry map[string]interface{}
		if err := json.Unmarshal(buf.Bytes(), &logEntry); err != nil {
			t.Fatalf("failed to parse log output: %v", err)
		}

		if logEntry["message"] != "user alice logged in at 12345" {
			t.Errorf("expected formatted message, got %v", logEntry["message"])
		}
	})

	t.Run("Send method", func(t *testing.T) {
		buf.Reset()
		logger.Info().Str("field", "value").Send()

		var logEntry map[string]interface{}
		if err := json.Unmarshal(buf.Bytes(), &logEntry); err != nil {
			t.Fatalf("failed to parse log output: %v", err)
		}

		if logEntry["field"] != "value" {
			t.Errorf("expected field value, got %v", logEntry["field"])
		}
		if _, exists := logEntry["message"]; exists {
			t.Errorf("expected no message field, but got one")
		}
	})
}

func TestEventMethodChaining(t *testing.T) {
	var buf bytes.Buffer
	traceID := "test-trace-chain"

	logger := &ContextLogger{
		logger:  zerolog.New(&buf).With().Timestamp().Logger(),
		traceID: traceID,
	}

	buf.Reset()
	logger.Info().
		Str("user", "bob").
		Int("age", 30).
		Bool("active", true).
		Msg("user profile")

	var logEntry map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &logEntry); err != nil {
		t.Fatalf("failed to parse log output: %v", err)
	}

	if logEntry["trace_id"] != traceID {
		t.Errorf("expected trace_id %s, got %v", traceID, logEntry["trace_id"])
	}
	if logEntry["user"] != "bob" {
		t.Errorf("expected user bob, got %v", logEntry["user"])
	}
	if logEntry["age"].(float64) != 30 {
		t.Errorf("expected age 30, got %v", logEntry["age"])
	}
	if logEntry["active"] != true {
		t.Errorf("expected active true, got %v", logEntry["active"])
	}
	if logEntry["message"] != "user profile" {
		t.Errorf("expected message 'user profile', got %v", logEntry["message"])
	}
}
