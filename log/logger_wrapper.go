package log

import (
    "log"
    cosmossdklog "cosmossdk.io/log"
)

// LoggerWrapper wraps a standard log.Logger to implement the cosmossdk.io/log.Logger interface.
type LoggerWrapper struct {
    *log.Logger
}

// Debug logs a debug message.
func (l *LoggerWrapper) Debug(msg string, keyvals ...interface{}) {
    l.Printf("DEBUG: "+msg, keyvals...)
}

// Info logs an info message.
func (l *LoggerWrapper) Info(msg string, keyvals ...interface{}) {
    l.Printf("INFO: "+msg, keyvals...)
}

// Error logs an error message.
func (l *LoggerWrapper) Error(msg string, keyvals ...interface{}) {
    l.Printf("ERROR: "+msg, keyvals...)
}

// Warn logs a warning message.
func (l *LoggerWrapper) Warn(msg string, keyvals ...interface{}) {
    l.Printf("WARN: "+msg, keyvals...)
}

// With returns a new logger with additional context (key-value pairs).
func (l *LoggerWrapper) With(keyvals ...interface{}) cosmossdklog.Logger {
    return &LoggerWrapper{Logger: l.Logger}
}

// Impl returns the underlying logger implementation.
func (l *LoggerWrapper) Impl() interface{} {
    return l.Logger
}

// NewLoggerWrapper creates a new LoggerWrapper
func NewLoggerWrapper(logger *log.Logger) *LoggerWrapper {
    return &LoggerWrapper{Logger: logger}
}
