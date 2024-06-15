package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel = zapcore.DebugLevel
	// InfoLevel is the default logging priority.
	InfoLevel = zapcore.InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel = zapcore.WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel = zapcore.ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel = zapcore.DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel = zapcore.PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel = zapcore.FatalLevel
)

var Log *zap.Logger

func init() {
	var err error

	// Configure log level
	logLevel := InfoLevel
	if level := os.Getenv("LOG_LEVEL"); level != "" {
		if err := logLevel.UnmarshalText([]byte(level)); err != nil {
			logLevel = InfoLevel
		}
	}

	// Create a production logger configuration
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(logLevel),
		Development:      false,
		Encoding:         "json", // Use JSON format for structured logging
		EncoderConfig:    getEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	Log, err = config.Build()
	if err != nil {
		panic(err)
	}
	defer Log.Sync() // Flushes buffer, if any
}

// getEncoderConfig returns the EncoderConfig for Zap logger
func getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder, // Use lowercase level encoding
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
}

// WithContext adds contextual information to the logger
func WithContext(fields ...zap.Field) *zap.Logger {
	return Log.With(fields...)
}

// Fatal logs a message, then calls os.Exit(1).
func Fatal(msg string, fields ...zap.Field) {
	Log.Fatal(msg, fields...)
}

// Panic logs a message, then panics.
func Panic(msg string, fields ...zap.Field) {
	Log.Panic(msg, fields...)
}

// Error logs a message
func Error(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}

// Info logs a message
func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

// Debug logs a message
func Debug(msg string, fields ...zap.Field) {
	Log.Debug(msg, fields...)
}

// Warn logs a message
func Warn(msg string, fields ...zap.Field) {
	Log.Warn(msg, fields...)
}

// DPanic logs a message, then panics.
func DPanic(msg string, fields ...zap.Field) {
	Log.DPanic(msg, fields...)
}
