package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

// InitLogger initializes the logger with custom settings such as time format and console output.
// It sets the global log level to Debug by default.
func InitLogger() {
	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "01/02/2006 03:04:05 PM"})

	// Set default global log level (can adjust as needed)
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}

// Info logs informational messages with optional structured fields (key-value pairs).
// Example usage: Info("Service started", map[string]any{"service": "auth", "status": "OK"})
func Info(msg string, fields map[string]any) {
	event := log.Info().Timestamp()
	for key, value := range fields {
		event.Interface(key, value)
	}
	event.Msg(msg)
}

// Error logs error messages along with optional structured fields.
// Example usage: Error(err, map[string]any{"action": "db_connect", "status": "failed"})
func Error(err error, fields map[string]any) {
	event := log.Error().Err(err).Timestamp()
	for key, value := range fields {
		event.Interface(key, value)
	}
	event.Msg("")
}

// Warn logs warning messages with optional structured fields.
// Example usage: Warn("Low disk space", map[string]any{"available": "500MB"})
func Warn(msg string, fields map[string]any) {
	event := log.Warn().Timestamp()
	for key, value := range fields {
		event.Interface(key, value)
	}
	event.Msg(msg)
}

// WithDuration logs informational messages with an additional "duration" field to track performance.
// Example usage: WithDuration("DB query executed", time.Duration(125 * time.Millisecond), map[string]any{"query": "SELECT * FROM users"})
func WithDuration(msg string, duration time.Duration, fields map[string]any) {
	event := log.Info().Timestamp().Str("duration", duration.String())
	for key, value := range fields {
		event.Interface(key, value)
	}
	event.Msg(msg)
}

// Debug logs debug messages with optional structured fields. Useful for tracing.
// Example usage: Debug("Cache miss", map[string]any{"key": "user_123"})
func Debug(msg string, fields map[string]any) {
	event := log.Debug().Timestamp()
	for key, value := range fields {
		event.Interface(key, value)
	}
	event.Msg(msg)
}
