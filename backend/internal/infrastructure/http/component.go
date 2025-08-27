package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

// Component provides lifecycle-managed HTTP server
type Component struct {
	server     *http.Server
	logger     *slog.Logger
	handler    http.Handler
	addr       string
	
	// Server state tracking
	serverStarted chan struct{}
	serverError   chan error
	
	// Configuration
	readTimeout    time.Duration
	writeTimeout   time.Duration
	idleTimeout    time.Duration
	headerTimeout  time.Duration
}

// ComponentOptions configures the HTTP component
type ComponentOptions struct {
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	IdleTimeout    time.Duration
	HeaderTimeout  time.Duration
}

// NewComponent creates a new HTTP server component
func NewComponent(addr string, handler http.Handler, logger *slog.Logger, opts *ComponentOptions) *Component {
	if opts == nil {
		opts = &ComponentOptions{
			ReadTimeout:   30 * time.Second,
			WriteTimeout:  30 * time.Second,
			IdleTimeout:   120 * time.Second,
			HeaderTimeout: 10 * time.Second,
		}
	}

	return &Component{
		addr:           addr,
		handler:        handler,
		logger:         logger,
		serverStarted:  make(chan struct{}),
		serverError:    make(chan error, 1),
		readTimeout:    opts.ReadTimeout,
		writeTimeout:   opts.WriteTimeout,
		idleTimeout:    opts.IdleTimeout,
		headerTimeout:  opts.HeaderTimeout,
	}
}

// Name returns the component name
func (c *Component) Name() string {
	return "http-server"
}

// Dependencies returns component dependencies
func (c *Component) Dependencies() []string {
	return []string{"database", "monitoring"}
}

// Start initializes and starts the HTTP server
func (c *Component) Start(ctx context.Context) error {
	c.server = &http.Server{
		Addr:    c.addr,
		Handler: c.handler,
		
		// Timeouts for graceful shutdown and security
		ReadTimeout:       c.readTimeout,
		WriteTimeout:      c.writeTimeout,
		IdleTimeout:       c.idleTimeout,
		ReadHeaderTimeout: c.headerTimeout,
	}
	
	// Start server in goroutine
	go func() {
		defer close(c.serverStarted)
		
		c.logger.Info("starting HTTP server", 
			slog.String("addr", c.addr),
			slog.Duration("read_timeout", c.readTimeout),
			slog.Duration("write_timeout", c.writeTimeout))
		
		if err := c.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			c.logger.Error("HTTP server failed", slog.Any("error", err))
			select {
			case c.serverError <- err:
			default:
			}
		}
	}()
	
	// Wait for server to start or fail with timeout
	startCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	
	select {
	case err := <-c.serverError:
		return fmt.Errorf("HTTP server failed to start: %w", err)
	case <-c.serverStarted:
		c.logger.Info("HTTP server started successfully")
		return nil
	case <-startCtx.Done():
		return fmt.Errorf("HTTP server start timeout: %w", startCtx.Err())
	}
}

// Stop gracefully shuts down the HTTP server
func (c *Component) Stop(ctx context.Context) error {
	if c.server == nil {
		return nil
	}
	
	c.logger.Info("starting HTTP server graceful shutdown")
	
	// Create shutdown context with timeout (leave 5s buffer from parent context)
	shutdownCtx, cancel := context.WithTimeout(ctx, 25*time.Second)
	defer cancel()
	
	// Shutdown server gracefully
	if err := c.server.Shutdown(shutdownCtx); err != nil {
		c.logger.Error("HTTP server failed to shutdown gracefully", 
			slog.Any("error", err))
		
		// Force close if graceful shutdown failed
		if closeErr := c.server.Close(); closeErr != nil {
			c.logger.Error("failed to force close HTTP server", 
				slog.Any("error", closeErr))
		}
		return fmt.Errorf("HTTP server shutdown failed: %w", err)
	}
	
	c.logger.Info("HTTP server stopped gracefully")
	return nil
}

// Server returns the underlying HTTP server
func (c *Component) Server() *http.Server {
	return c.server
}

// HealthCheck verifies HTTP server health
func (c *Component) HealthCheck(ctx context.Context) error {
	if c.server == nil {
		return fmt.Errorf("server not initialized")
	}
	
	// Create a simple health check request
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	
	// Try to connect to the health endpoint
	healthURL := fmt.Sprintf("http://%s/health", c.addr)
	req, err := http.NewRequestWithContext(ctx, "GET", healthURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create health check request: %w", err)
	}
	
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("health check request failed: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("health check failed with status: %d", resp.StatusCode)
	}
	
	return nil
}