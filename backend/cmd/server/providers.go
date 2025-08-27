package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/config"
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/persistence/postgres"
	db "github.com/fernandobandeira/djinn/backend/internal/infrastructure/persistence/postgres/generated"
	"github.com/fernandobandeira/djinn/backend/internal/graph/resolver"
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/server"
	"github.com/fernandobandeira/djinn/backend/internal/observability"
)

// ProvideConfig loads configuration from environment
func ProvideConfig() (*config.Config, error) {
	return config.Load()
}

// ProvideQueries creates database queries from the database connection
func ProvideQueries(database *postgres.DB) *db.Queries {
	if database == nil {
		panic("database is required")
	}
	if database.Queries == nil {
		panic("database queries not initialized")
	}
	return database.Queries
}


// ProvideDatabase creates a database connection
func ProvideDatabase(cfg *config.Config, logger *slog.Logger) (*postgres.DB, error) {
	// Validate input parameters
	if cfg == nil {
		return nil, fmt.Errorf("config is required")
	}
	if logger == nil {
		return nil, fmt.Errorf("logger is required")
	}
	
	// Use PgBouncer URL if available, otherwise direct database URL
	url := cfg.PgBouncerURL
	if url == "" {
		url = cfg.DatabaseURL
	}
	if url == "" {
		return nil, fmt.Errorf("database URL is required")
	}
	
	db, err := postgres.Connect(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	
	logger.Info("Database connection established",
		"max_connections", 25,
		"min_connections", 5)
	
	return db, nil
}

// ProvideServer creates the HTTP server
func ProvideServer(cfg *config.Config, logger *slog.Logger, db *postgres.DB, resolver *resolver.Resolver) *server.Server {
	if cfg == nil {
		panic("config is required")
	}
	if logger == nil {
		panic("logger is required")
	}
	if db == nil {
		panic("database is required")
	}
	if resolver == nil {
		panic("resolver is required")
	}
	return server.NewServer(cfg, logger, db, resolver)
}

// ProvideMonitoring initializes the integrated monitoring service
func ProvideMonitoring(cfg *config.Config, logger *slog.Logger) (*observability.MonitoringService, func(), error) {
	if cfg == nil {
		return nil, nil, fmt.Errorf("config is required")
	}
	if logger == nil {
		return nil, nil, fmt.Errorf("logger is required")
	}
	
	monitoringConfig := &observability.MonitoringConfig{
		// Metrics
		MetricsEnabled: cfg.MetricsEnabled,
		MetricsPort:    cfg.MetricsPort,
		
		// Tracing
		TracingEnabled:    cfg.TracingEnabled,
		TracingSampleRate: cfg.TracingSampleRate,
		OTLPEndpoint:      cfg.OTELEndpoint,
		
		// Analytics
		AnalyticsEnabled: cfg.AnalyticsEnabled,
		PostHogAPIKey:    cfg.PostHogAPIKey,
		
		// Service info
		ServiceName:    cfg.ServiceName,
		ServiceVersion: cfg.ServiceVersion,
		Environment:    cfg.Environment,
	}
	
	// Get monitoring service with context-aware shutdown
	service, shutdownWithContext, err := observability.NewMonitoringService(context.Background(), monitoringConfig, logger)
	if err != nil {
		return nil, nil, err
	}
	
	// Wrap the context-aware shutdown for wire compatibility
	wrappedShutdown := func() {
		_ = shutdownWithContext(context.Background())
	}
	
	// Return the service and wrapped shutdown
	return service, wrappedShutdown, nil
}

// Legacy providers removed - now using lifecycle-managed application

