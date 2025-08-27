package monitoring

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/fernandobandeira/djinn/backend/internal/observability"
)

// Component provides lifecycle-managed monitoring service
type Component struct {
	service       *observability.MonitoringService
	shutdownFunc  func(context.Context) error
	logger        *slog.Logger
	config        *observability.MonitoringConfig
}

// NewComponent creates a new monitoring component
func NewComponent(config *observability.MonitoringConfig, logger *slog.Logger) (*Component, error) {
	if config == nil {
		// Create default monitoring config if none provided
		config = &observability.MonitoringConfig{
			ServiceName:    "djinn-api",
			ServiceVersion: "1.0.0",
			Environment:    "development",
			MetricsEnabled: true,
			TracingEnabled: true,
		}
	}

	// Create monitoring service
	ctx := context.Background()
	service, shutdownFunc, err := observability.NewMonitoringService(ctx, config, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create monitoring service: %w", err)
	}
	
	return &Component{
		service:      service,
		shutdownFunc: shutdownFunc,
		logger:       logger,
		config:       config,
	}, nil
}

// Name returns the component name
func (c *Component) Name() string {
	return "monitoring"
}

// Dependencies returns component dependencies (monitoring has no dependencies)
func (c *Component) Dependencies() []string {
	return []string{}
}

// Start initializes the monitoring service
func (c *Component) Start(ctx context.Context) error {
	c.logger.Info("starting monitoring component",
		slog.String("service", c.config.ServiceName),
		slog.Bool("metrics_enabled", c.config.MetricsEnabled),
		slog.Bool("tracing_enabled", c.config.TracingEnabled))
	
	// Monitoring service is typically passive, just verify it's ready
	if c.service != nil {
		// Test basic functionality
		c.logger.Debug("monitoring service initialized successfully")
	}
	
	c.logger.Info("monitoring component started successfully")
	return nil
}

// Stop gracefully shuts down the monitoring service
func (c *Component) Stop(ctx context.Context) error {
	c.logger.Info("stopping monitoring component")
	
	if c.shutdownFunc != nil {
		// Use timeout context for shutdown - give monitoring time to flush
		shutdownCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
		defer cancel()
		
		if err := c.shutdownFunc(shutdownCtx); err != nil {
			c.logger.Error("monitoring shutdown failed", slog.Any("error", err))
			return fmt.Errorf("monitoring shutdown failed: %w", err)
		}
	}
	
	c.logger.Info("monitoring component stopped successfully")
	return nil
}

// Service returns the underlying monitoring service
func (c *Component) Service() *observability.MonitoringService {
	return c.service
}

// HealthCheck verifies monitoring service health
func (c *Component) HealthCheck(ctx context.Context) error {
	if c.service == nil {
		return fmt.Errorf("monitoring service not initialized")
	}
	
	// Basic health check - monitoring service should be responsive
	// This could be expanded to check individual metrics/tracing endpoints
	return nil
}