package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/fernandobandeira/djinn/backend/internal/graph/resolver"
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/config"
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/database"
	httpComponent "github.com/fernandobandeira/djinn/backend/internal/infrastructure/http"
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/lifecycle"
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/monitoring"
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/persistence/postgres"
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/server"
	"github.com/fernandobandeira/djinn/backend/internal/observability"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Application represents the main application with lifecycle management
type Application struct {
	config             *config.Config
	logger             *slog.Logger
	lifecycleManager   *lifecycle.Manager
	
	// Component references for external access
	dbComponent       *database.Component
	monitoringComponent *monitoring.Component
	httpComponent     *httpComponent.Component
}

// NewApplication creates a new application instance with lifecycle management
func NewApplication(cfg *config.Config, logger *slog.Logger) (*Application, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config is required")
	}
	if logger == nil {
		return nil, fmt.Errorf("logger is required")
	}

	// Create lifecycle manager
	manager := lifecycle.NewManager(logger)
	
	app := &Application{
		config:           cfg,
		logger:           logger,
		lifecycleManager: manager,
	}
	
	// Initialize all components
	if err := app.initializeComponents(); err != nil {
		return nil, fmt.Errorf("failed to initialize components: %w", err)
	}
	
	return app, nil
}

// initializeComponents sets up all system components with proper dependencies
func (app *Application) initializeComponents() error {
	// 1. Initialize monitoring component (no dependencies)
	monitoringConfig := &observability.MonitoringConfig{
		MetricsEnabled:    app.config.MetricsEnabled,
		MetricsPort:       app.config.MetricsPort,
		TracingEnabled:    app.config.TracingEnabled,
		TracingSampleRate: app.config.TracingSampleRate,
		OTLPEndpoint:      app.config.OTELEndpoint,
		AnalyticsEnabled:  app.config.AnalyticsEnabled,
		PostHogAPIKey:     app.config.PostHogAPIKey,
		ServiceName:       app.config.ServiceName,
		ServiceVersion:    app.config.ServiceVersion,
		Environment:       app.config.Environment,
	}
	
	var err error
	app.monitoringComponent, err = monitoring.NewComponent(monitoringConfig, app.logger)
	if err != nil {
		return fmt.Errorf("failed to create monitoring component: %w", err)
	}
	
	if err := app.lifecycleManager.RegisterComponent(app.monitoringComponent); err != nil {
		return fmt.Errorf("failed to register monitoring component: %w", err)
	}
	
	// 2. Initialize database component (no dependencies)
	dbURL := app.config.PgBouncerURL
	if dbURL == "" {
		dbURL = app.config.DatabaseURL
	}
	if dbURL == "" {
		return fmt.Errorf("database URL is required")
	}
	
	// Create pgxpool config
	poolConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return fmt.Errorf("failed to parse database URL: %w", err)
	}
	
	// Configure connection pool
	poolConfig.MaxConns = 25
	poolConfig.MinConns = 5
	
	app.dbComponent = database.NewComponent(poolConfig, app.logger)
	if err := app.lifecycleManager.RegisterComponent(app.dbComponent); err != nil {
		return fmt.Errorf("failed to register database component: %w", err)
	}
	
	// 3. Initialize HTTP server component (depends on database and monitoring)
	// Create the HTTP handler using the original server setup
	handler, err := app.createHTTPHandler()
	if err != nil {
		return fmt.Errorf("failed to create HTTP handler: %w", err)
	}
	
	httpOpts := &httpComponent.ComponentOptions{
		ReadTimeout:   app.config.ReadTimeout,
		WriteTimeout:  app.config.WriteTimeout,
		IdleTimeout:   app.config.IdleTimeout,
		HeaderTimeout: 10 * time.Second,
	}
	
	addr := fmt.Sprintf(":%d", app.config.Port)
	app.httpComponent = httpComponent.NewComponent(addr, handler, app.logger, httpOpts)
	
	if err := app.lifecycleManager.RegisterComponent(app.httpComponent); err != nil {
		return fmt.Errorf("failed to register HTTP component: %w", err)
	}
	
	return nil
}

// createHTTPHandler creates the HTTP handler using the existing server infrastructure
func (app *Application) createHTTPHandler() (http.Handler, error) {
	// Create a postgres.DB wrapper for compatibility with existing server code
	// This bridges the new component-based DB with the existing server expectations
	db := &postgres.DB{}
	db.SetPool(app.dbComponent.Pool())
	
	// Create resolver and server using existing patterns
	resolver := app.createResolver(db)
	srv := server.NewServer(app.config, app.logger, db, resolver)
	
	// Return the server's router as the handler
	return srv.Handler(), nil
}

// createResolver creates the GraphQL resolver with minimal setup
func (app *Application) createResolver(db *postgres.DB) *resolver.Resolver {
	// For now, use the interface-based constructor for testing/minimal setup
	// This can be expanded later with full DI when all handlers are needed
	return resolver.NewResolverWithInterfaces(db.Queries, app.logger)
}

// Run starts the application and blocks until shutdown signal
func (app *Application) Run() error {
	app.logger.Info("starting application",
		slog.String("service", app.config.ServiceName),
		slog.String("version", app.config.ServiceVersion),
		slog.String("environment", app.config.Environment))
	
	// Run lifecycle manager (blocks until shutdown signal)
	if err := app.lifecycleManager.Run(); err != nil {
		app.logger.Error("application stopped with error", slog.Any("error", err))
		return err
	}
	
	app.logger.Info("application stopped gracefully")
	return nil
}

// GetDatabaseComponent returns the database component for external access
func (app *Application) GetDatabaseComponent() *database.Component {
	return app.dbComponent
}

// GetMonitoringComponent returns the monitoring component for external access
func (app *Application) GetMonitoringComponent() *monitoring.Component {
	return app.monitoringComponent
}

// GetHTTPComponent returns the HTTP component for external access
func (app *Application) GetHTTPComponent() *httpComponent.Component {
	return app.httpComponent
}