package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fernandobandeira/djinn/backend/internal/config"
	"github.com/fernandobandeira/djinn/backend/internal/database"
	"github.com/fernandobandeira/djinn/backend/internal/server"
)

// Application represents the main application with all dependencies
type Application struct {
	config          *config.Config
	logger          *slog.Logger
	db              *database.DB
	server          *server.Server
	shutdownTracing func(context.Context) error
}

// NewApplication creates a new application instance
func NewApplication(cfg *config.Config, logger *slog.Logger, db *database.DB, srv *server.Server) *Application {
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
	
	// Shutdown tracing if configured
	if app.shutdownTracing != nil {
		if err := app.shutdownTracing(ctx); err != nil {
			app.logger.Error("Failed to shutdown tracing", "error", err)
		}
	}
	
	return nil
}