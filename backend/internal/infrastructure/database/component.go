package database

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/persistence/postgres/generated"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Component provides lifecycle-managed database connectivity with graceful shutdown
type Component struct {
	pool           *pgxpool.Pool
	queries        *generated.Queries
	logger         *slog.Logger
	config         *pgxpool.Config
	
	// Transaction tracking for graceful shutdown
	activeTxMutex  sync.RWMutex
	activeTx       map[string]context.CancelFunc
	shutdownCtx    context.Context
	shutdownCancel context.CancelFunc
	isShuttingDown bool
}

// NewComponent creates a new database component
func NewComponent(config *pgxpool.Config, logger *slog.Logger) *Component {
	shutdownCtx, shutdownCancel := context.WithCancel(context.Background())
	
	return &Component{
		config:         config,
		logger:         logger,
		activeTx:       make(map[string]context.CancelFunc),
		shutdownCtx:    shutdownCtx,
		shutdownCancel: shutdownCancel,
	}
}

// Name returns the component name
func (c *Component) Name() string {
	return "database"
}

// Dependencies returns component dependencies (database has no dependencies)
func (c *Component) Dependencies() []string {
	return []string{}
}

// Start initializes the database connection pool
func (c *Component) Start(ctx context.Context) error {
	var err error
	c.pool, err = pgxpool.NewWithConfig(ctx, c.config)
	if err != nil {
		return fmt.Errorf("failed to create connection pool: %w", err)
	}
	
	// Test connection
	if err := c.pool.Ping(ctx); err != nil {
		c.pool.Close()
		return fmt.Errorf("failed to ping database: %w", err)
	}
	
	// Initialize queries
	c.queries = generated.New(c.pool)
	
	c.logger.Info("database component started",
		slog.Int("max_conns", int(c.config.MaxConns)),
		slog.Int("min_conns", int(c.config.MinConns)))
	
	return nil
}

// Stop gracefully shuts down the database connection pool
func (c *Component) Stop(ctx context.Context) error {
	c.logger.Info("starting database graceful shutdown")
	
	// Prevent new transactions
	c.activeTxMutex.Lock()
	c.isShuttingDown = true
	c.activeTxMutex.Unlock()
	
	// Signal shutdown to prevent new transactions
	c.shutdownCancel()
	
	// Wait for active transactions with timeout
	shutdownComplete := make(chan struct{})
	
	go func() {
		defer close(shutdownComplete)
		
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		
		for {
			c.activeTxMutex.RLock()
			activeTxCount := len(c.activeTx)
			c.activeTxMutex.RUnlock()
			
			if activeTxCount == 0 {
				c.logger.Info("all transactions completed")
				return
			}
			
			c.logger.Info("waiting for active transactions",
				slog.Int("count", activeTxCount))
			
			select {
			case <-ticker.C:
				continue
			case <-ctx.Done():
				c.logger.Warn("shutdown timeout reached, cancelling active transactions",
					slog.Int("count", activeTxCount))
				
				// Cancel remaining transactions
				c.activeTxMutex.Lock()
				for txID, cancel := range c.activeTx {
					c.logger.Warn("force cancelling transaction", slog.String("tx_id", txID))
					cancel()
				}
				c.activeTxMutex.Unlock()
				return
			}
		}
	}()
	
	select {
	case <-shutdownComplete:
		c.logger.Info("database transactions drained successfully")
	case <-ctx.Done():
		c.logger.Warn("database shutdown timeout reached")
	}
	
	// Close connection pool
	if c.pool != nil {
		c.pool.Close()
	}
	
	c.logger.Info("database component stopped")
	return nil
}

// Pool returns the underlying connection pool with safety checks
func (c *Component) Pool() *pgxpool.Pool {
	c.activeTxMutex.RLock()
	defer c.activeTxMutex.RUnlock()
	
	// SAFETY FIX: Check component state and nil pointers
	if c.isShuttingDown || c.pool == nil {
		return nil
	}
	return c.pool
}

// Queries returns the generated database queries with safety checks
func (c *Component) Queries() *generated.Queries {
	c.activeTxMutex.RLock()
	defer c.activeTxMutex.RUnlock()
	
	// SAFETY FIX: Check component state and nil pointers
	if c.isShuttingDown || c.queries == nil {
		return nil
	}
	return c.queries
}

// BeginTx starts a tracked transaction for graceful shutdown
func (c *Component) BeginTx(ctx context.Context) (pgx.Tx, error) {
	// Atomic shutdown check and transaction registration - FIXES TOCTOU
	c.activeTxMutex.Lock()
	
	// Check shutdown state atomically while holding the lock
	if c.isShuttingDown {
		c.activeTxMutex.Unlock()
		return nil, fmt.Errorf("database is shutting down")
	}
	
	// Double-check with context
	select {
	case <-c.shutdownCtx.Done():
		c.activeTxMutex.Unlock()
		return nil, fmt.Errorf("database is shutting down")
	default:
	}
	
	// Create transaction context that can be cancelled during shutdown
	txCtx, cancel := context.WithCancel(ctx)
	txID := uuid.New().String()
	
	// Pre-register transaction in tracking while holding lock
	c.activeTx[txID] = cancel
	
	// Release lock for database call to avoid holding lock during I/O
	c.activeTxMutex.Unlock()
	
	// Attempt to create the actual transaction
	tx, err := c.pool.BeginTx(txCtx, pgx.TxOptions{})
	
	// Re-acquire lock to clean up if needed
	if err != nil {
		c.activeTxMutex.Lock()
		delete(c.activeTx, txID)
		c.activeTxMutex.Unlock()
		cancel()
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	
	// Transaction created successfully, return wrapped transaction
	return &trackedTx{
		Tx:     tx,
		id:     txID,
		cancel: cancel,
		parent: c,
	}, nil
}

// trackedTx wraps pgx.Tx to provide cleanup on completion
type trackedTx struct {
	pgx.Tx
	id     string
	cancel context.CancelFunc
	parent *Component
}

// Commit commits the transaction and cleans up tracking
func (t *trackedTx) Commit(ctx context.Context) error {
	defer t.cleanup()
	return t.Tx.Commit(ctx)
}

// Rollback rolls back the transaction and cleans up tracking
func (t *trackedTx) Rollback(ctx context.Context) error {
	defer t.cleanup()
	return t.Tx.Rollback(ctx)
}

// cleanup removes transaction from active tracking
func (t *trackedTx) cleanup() {
	t.cancel()
	t.parent.activeTxMutex.Lock()
	delete(t.parent.activeTx, t.id)
	t.parent.activeTxMutex.Unlock()
}

// HealthCheck verifies database connectivity
func (c *Component) HealthCheck(ctx context.Context) error {
	if c.pool == nil {
		return fmt.Errorf("database pool not initialized")
	}
	return c.pool.Ping(ctx)
}

// GetActiveTransactionCount returns the number of active transactions
func (c *Component) GetActiveTransactionCount() int {
	c.activeTxMutex.RLock()
	defer c.activeTxMutex.RUnlock()
	return len(c.activeTx)
}

// Close provides postgres.DB compatibility (delegates to component lifecycle)
func (c *Component) Close() {
	// Component closure is handled by lifecycle manager
	// This method exists for postgres.DB compatibility
}

// Ping provides postgres.DB compatibility
func (c *Component) Ping(ctx context.Context) error {
	return c.HealthCheck(ctx)
}

// GetPool provides postgres.DB compatibility
func (c *Component) GetPool() *pgxpool.Pool {
	return c.pool
}

// Health provides postgres.DB compatibility
func (c *Component) Health(ctx context.Context) error {
	return c.HealthCheck(ctx)
}