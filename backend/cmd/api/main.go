package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/yourusername/djinn/backend/internal/config"
	"github.com/yourusername/djinn/backend/internal/database"
	"github.com/yourusername/djinn/backend/internal/middleware"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	// Setup logger
	logger := setupLogger(cfg)
	logger.Info("Starting Djinn API server",
		"version", cfg.ServiceVersion,
		"environment", cfg.Environment,
		"port", cfg.Port,
	)

	// Connect to database
	db, err := database.Connect(cfg.PgBouncerURL)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	// Setup router
	router := setupRouter(cfg, logger, db)

	// Setup HTTP server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	// Start server in a goroutine
	go func() {
		logger.Info("HTTP server starting", "address", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Failed to start HTTP server", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown", "error", err)
	}

	logger.Info("Server shutdown complete")
}

func setupLogger(cfg *config.Config) *slog.Logger {
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

func setupRouter(cfg *config.Config, logger *slog.Logger, db *database.DB) *chi.Mux {
	r := chi.NewRouter()

	// Global middleware
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger(logger))
	r.Use(middleware.Recovery(logger))
	r.Use(middleware.SecurityHeaders)
	r.Use(middleware.CORS([]string{"http://localhost:3000", "http://localhost:4200"})) // Add your frontend URLs
	r.Use(chimiddleware.Compress(5))
	r.Use(chimiddleware.Timeout(60 * time.Second))

	// Health check endpoints
	r.Get("/health", healthCheckHandler)
	r.Get("/ready", readinessHandler(db))

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.ContentType("application/json"))
		
		// GraphQL endpoint will be added in the next task
		// r.Handle("/graphql", graphqlHandler)
		
		// Placeholder for REST endpoints if needed
		r.Get("/version", versionHandler(cfg))
	})

	// GraphQL playground (development only)
	if cfg.EnablePlayground {
		// Will be added with gqlgen in the next task
		// r.Get("/playground", playgroundHandler)
	}

	// Metrics endpoint
	if cfg.MetricsEnabled {
		// Will be added with Prometheus metrics
		// r.Handle("/metrics", promhttp.Handler())
	}

	return r
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"healthy"}`))
}

func readinessHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		if err := db.Ping(ctx); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"status":"not_ready","error":"database_unavailable"}`))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ready"}`))
	}
}

func versionHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"service":"%s","version":"%s","environment":"%s"}`,
			cfg.ServiceName, cfg.ServiceVersion, cfg.Environment)
	}
}