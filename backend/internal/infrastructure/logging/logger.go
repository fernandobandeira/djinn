package logging

import (
	"log/slog"
	"os"
	
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/config"
)

// NewLogger creates a new structured logger based on configuration
func NewLogger(cfg *config.Config) *slog.Logger {
	var handler slog.Handler
	
	opts := &slog.HandlerOptions{
		Level: parseLogLevel(cfg.LogLevel),
		AddSource: cfg.Environment == "development",
	}
	
	// Use JSON handler if configured or in production
	if cfg.LogJSON || cfg.Environment == "production" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}
	
	logger := slog.New(handler)
	
	// Set as default logger
	slog.SetDefault(logger)
	
	logger.Info("Logger initialized",
		"environment", cfg.Environment,
		"level", cfg.LogLevel,
	)
	
	return logger
}

// parseLogLevel converts string log level to slog.Level
func parseLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}