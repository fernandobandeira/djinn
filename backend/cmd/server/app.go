package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/config"
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/persistence/postgres"
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/server"
	"github.com/fernandobandeira/djinn/backend/internal/observability"
)

// Application represents the main application with all dependencies
type Application struct {
	config             *config.Config
	logger             *slog.Logger
	db                 *postgres.DB
	server             *server.Server
	monitoring         *observability.MonitoringService
	shutdownMonitoring func(context.Context) error
}

// NewApplication creates a new application instance
func NewApplication(cfg *config.Config, logger *slog.Logger, db *postgres.DB, srv *server.Server) *Application {
	return &Application{
		config: cfg,
		logger: logger,
		db:     db,
		server: srv,
	}
}

// Run starts the application and handles graceful shutdown
func (app *Application) Run() error {
	// Start server in a goroutine
	errCh := make(chan error, 1)
	go func() {
		if err := app.server.Start(); err != nil {
			errCh <- err
		}
	}()
	
	// Wait for interrupt signal or server error
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	
	select {
	case err := <-errCh:
		return err
	case <-quit:
		app.logger.Info("Interrupt signal received")
	}
	
	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	if err := app.server.Shutdown(ctx); err != nil {
		app.logger.Error("Server forced to shutdown", "error", err)
		return err
	}
	
	// Close database connection
	app.db.Close()
	
	// Shutdown monitoring if configured
	if app.shutdownMonitoring != nil {
		if err := app.shutdownMonitoring(ctx); err != nil {
			app.logger.Error("Failed to shutdown monitoring", "error", err)
		}
	}
	
	return nil
}