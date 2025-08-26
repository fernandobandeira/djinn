package logger

import (
	"log/slog"
	"os"

	"github.com/fernandobandeira/djinn/backend/internal/config"
)

// New creates a new slog logger based on configuration
func New(cfg *config.Config) *slog.Logger {
	var level slog.Level
	switch cfg.LogLevel {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}
	
	opts := &slog.HandlerOptions{
		Level: level,
	}
	
	var handler slog.Handler
	if cfg.LogJSON {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}
	
	return slog.New(handler)
}