package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/fernandobandeira/djinn/backend/internal/database"
	"github.com/fernandobandeira/djinn/backend/internal/graph/resolver"
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/config"
	"github.com/fernandobandeira/djinn/backend/internal/middleware"
)

// Server represents the HTTP server and its dependencies
type Server struct {
	config   *config.Config
	logger   *slog.Logger
	db       *database.DB
	resolver *resolver.Resolver  // GraphQL resolver
	router   *chi.Mux
	httpSrv  *http.Server
}

// NewServer creates a new server instance
func NewServer(cfg *config.Config, logger *slog.Logger, db *database.DB, res *resolver.Resolver) *Server {
	s := &Server{
		config:   cfg,
		logger:   logger,
		db:       db,
		resolver: res,
	}
	
	s.setupRouter()
	s.httpSrv = &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      s.router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
	
	return s
}

func (s *Server) setupRouter() {
	s.router = chi.NewRouter()
	
	// Global middleware
	s.router.Use(chimiddleware.RealIP)
	s.router.Use(chimiddleware.Recoverer)
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.Logger(s.logger))
	s.router.Use(middleware.Recovery(s.logger))
	s.router.Use(middleware.SecurityHeaders)
	s.router.Use(middleware.CORS([]string{"http://localhost:3000", "http://localhost:4200"}))
	s.router.Use(chimiddleware.Compress(5))
	s.router.Use(chimiddleware.Timeout(60 * time.Second))
	
	// Add OpenTelemetry HTTP instrumentation if enabled
	if s.config.TracingEnabled {
		s.router.Use(middleware.OTelHTTPMiddleware(s.config.ServiceName))
	}
	
	// Mount routes
	s.mountRoutes()
}

func (s *Server) mountRoutes() {
	// Health check endpoints
	s.router.Get("/health", s.healthHandler)
	s.router.Get("/ready", s.readinessHandler)
	
	// API routes
	s.router.Route("/api", func(r chi.Router) {
		// GraphQL endpoint
		r.Handle("/graphql", s.graphqlHandler())
		
		// Version endpoint
		r.Get("/version", s.versionHandler)
	})
	
	// GraphQL playground (development only)
	if s.config.EnablePlayground {
		s.router.Handle("/playground", s.playgroundHandler())
	}
	
	// Metrics endpoint
	if s.config.MetricsEnabled {
		// Will be added with Prometheus
		// s.router.Handle("/metrics", promhttp.Handler())
	}
}

// Start starts the HTTP server
func (s *Server) Start() error {
	s.logger.Info("HTTP server starting", 
		"address", s.httpSrv.Addr,
		"version", s.config.ServiceVersion,
		"environment", s.config.Environment,
	)
	
	if err := s.httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to start server: %w", err)
	}
	
	return nil
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down server...")
	
	if err := s.httpSrv.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %w", err)
	}
	
	s.logger.Info("Server shutdown complete")
	return nil
}