package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/fernandobandeira/djinn/backend/internal/config"
	"github.com/fernandobandeira/djinn/backend/internal/database"
	"github.com/fernandobandeira/djinn/backend/internal/observability"
	"github.com/fernandobandeira/djinn/backend/internal/server"
)

// ProvideConfig loads configuration from environment
func ProvideConfig() (*config.Config, error) {
	return config.Load()
}

// ProvideLogger creates a logger instance
func ProvideLogger(cfg *config.Config) *slog.Logger {
	var handler slog.Handler
	
	opts := &slog.HandlerOptions{
		Level: parseLogLevel(cfg.LogLevel),
	}
	
	if cfg.LogJSON {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}
	
	logger := slog.New(handler)
	slog.SetDefault(logger)
	
	return logger
}

// ProvideDatabase creates a database connection
func ProvideDatabase(cfg *config.Config, logger *slog.Logger) (*database.DB, error) {
	// Use PgBouncer URL if available, otherwise direct database URL
	url := cfg.PgBouncerURL
	if url == "" {
		url = cfg.DatabaseURL
	}
	
	db, err := database.Connect(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	
	logger.Info("Database connection established",
		"max_connections", 25,
		"min_connections", 5)
	
	return db, nil
}

// ProvideServer creates the HTTP server
func ProvideServer(cfg *config.Config, logger *slog.Logger, db *database.DB) *server.Server {
	return server.NewServer(cfg, logger, db)
}

// ProvideTracing initializes OpenTelemetry tracing
func ProvideTracing(cfg *config.Config) (func(context.Context) error, error) {
	tracingConfig := observability.TracingConfig{
		ServiceName:    cfg.ServiceName,
		ServiceVersion: cfg.ServiceVersion,
		Environment:    cfg.Environment,
		OTLPEndpoint:   cfg.OTELEndpoint,
		Enabled:        cfg.TracingEnabled,
		SampleRate:     cfg.TracingSampleRate,
	}
	
	return observability.InitTracer(context.Background(), tracingConfig)
}

// ProvideApplication creates the main application with tracing
func ProvideApplication(
	cfg *config.Config,
	logger *slog.Logger,
	db *database.DB,
	srv *server.Server,
	shutdownTracing func(context.Context) error,
) *Application {
	return &Application{
		config:          cfg,
		logger:          logger,
		db:              db,
		server:          srv,
		shutdownTracing: shutdownTracing,
	}
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