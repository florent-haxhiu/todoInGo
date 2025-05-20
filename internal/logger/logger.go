package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
)

var globalLogger *slog.Logger

const (
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
)

type LoggerOptions struct {
	Level     slog.Level
	Format    string
	AddSource bool
	Output    io.Writer
}

func DefaultOptions() LoggerOptions {
	return LoggerOptions{
		Level:     LevelInfo,
		Format:    "text",
		AddSource: false,
		Output:    os.Stdout,
	}
}

func Initialize(opts LoggerOptions) {
	var handler slog.Handler

	switch opts.Format {
	case "json":
		handler = slog.NewJSONHandler(opts.Output, &slog.HandlerOptions{
			Level:     opts.Level,
			AddSource: opts.AddSource,
		})
	default:
		handler = slog.NewTextHandler(opts.Output, &slog.HandlerOptions{
			Level:     opts.Level,
			AddSource: opts.AddSource,
		})
	}

	globalLogger = slog.New(handler)
	slog.SetDefault(globalLogger)
}

func InitializeDefault() {
	opts := DefaultOptions()

	env := os.Getenv("ENVIRONMENT")
	switch env {
	case "production":
		opts.Format = "json"
		opts.Level = LevelInfo
	case "development":
		opts.Format = "json"
		opts.Level = LevelDebug
		opts.AddSource = true
	}

	Initialize(opts)

	globalLogger.Debug("Logger initialized",
		"environment", env,
		"format", opts.Format,
		"level", opts.Level.String())
}

func WithRequestID(requestID string) *slog.Logger {
	return globalLogger.With("request_id", requestID)
}

type loggerKeyType struct{}

var loggerKey = loggerKeyType{}

func NewContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func FromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerKey).(*slog.Logger); ok {
		return logger
	}
	return globalLogger
}

func Debug() *slog.Logger {
	return globalLogger.With(slog.String("level", "DEBUG"))
}

func Info() *slog.Logger {
	return globalLogger.With(slog.String("level", "INFO"))
}

func Warn() *slog.Logger {
	return globalLogger.With(slog.String("level", "WARN"))
}

func Error() *slog.Logger {
	return globalLogger.With(slog.String("level", "ERROR"))
}

func DebugMsg(msg string, args ...any) {
	globalLogger.Debug(msg, args...)
}

func InfoMsg(msg string, args ...any) {
	globalLogger.Info(msg, args...)
}

func WarnMsg(msg string, args ...any) {
	globalLogger.Warn(msg, args...)
}

func ErrorMsg(msg string, args ...any) {
	globalLogger.Error(msg, args...)
}
