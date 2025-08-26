package database

import (
	"context"
	"fmt"
	"time"

	"github.com/exaring/otelpgx"
	"github.com/fernandobandeira/djinn/backend/internal/database/generated"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Config holds database configuration
type Config struct {
	URL             string
	PgBouncerURL    string
	MaxConnections  int32
	MinConnections  int32
	MaxConnLifetime time.Duration
}

// DB wraps the database connection and queries
type DB struct {
	pool    *pgxpool.Pool
	Queries *generated.Queries
}

// NewConnection creates a new database connection
func NewConnection(config Config) (*DB, error) {
	// Parse connection URL and create pool config
	poolConfig, err := pgxpool.ParseConfig(config.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database URL: %w", err)
	}

	// Configure connection pool
	poolConfig.MaxConns = config.MaxConnections
	poolConfig.MinConns = config.MinConnections
	poolConfig.MaxConnLifetime = config.MaxConnLifetime
	
	// Add OpenTelemetry instrumentation to pgx
	poolConfig.ConnConfig.Tracer = otelpgx.NewTracer()

	// Create database connection pool
	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create database pool: %w", err)
	}

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Create queries instance - pgx/v5 uses the pool directly
	queries := generated.New(pool)

	return &DB{
		pool:    pool,
		Queries: queries,
	}, nil
}

// NewPgBouncerConnection creates a connection through PgBouncer
func NewPgBouncerConnection(config Config) (*DB, error) {
	if config.PgBouncerURL == "" {
		return NewConnection(config)
	}

	// Use PgBouncer URL instead of direct connection
	pgBouncerConfig := config
	pgBouncerConfig.URL = config.PgBouncerURL
	
	return NewConnection(pgBouncerConfig)
}

// Close closes the database connection pool
func (d *DB) Close() {
	if d.pool != nil {
		d.pool.Close()
	}
}

// Health checks database connectivity
func (d *DB) Health(ctx context.Context) error {
	return d.pool.Ping(ctx)
}

// GetPool returns the underlying pgxpool.Pool for advanced use cases
func (d *DB) GetPool() *pgxpool.Pool {
	return d.pool
}

// Ping checks database connectivity
func (d *DB) Ping(ctx context.Context) error {
	return d.pool.Ping(ctx)
}

// Connect creates a database connection from a URL string
func Connect(url string) (*DB, error) {
	config := Config{
		URL:             url,
		MaxConnections:  25,
		MinConnections:  5,
		MaxConnLifetime: 1 * time.Hour,
	}
	return NewConnection(config)
}