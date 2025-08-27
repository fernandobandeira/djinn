package database

import (
	"context"
	"log/slog"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	postgrescontainer "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/goleak"
)

// Common goleak ignore patterns for testcontainers and network operations
var testcontainersIgnorePatterns = []goleak.Option{
	goleak.IgnoreTopFunction("github.com/testcontainers/testcontainers-go.(*Reaper).connect.func1"),
	goleak.IgnoreTopFunction("net.doBlockingWithCtx.func1"),
	goleak.IgnoreTopFunction("net._C2func_getaddrinfo"),
	goleak.IgnoreTopFunction("github.com/jackc/pgx/v5/pgxpool.(*Pool).createIdleResources"),
}

// TestDatabaseComponent_GracefulShutdown tests the database component lifecycle with graceful shutdown
func TestDatabaseComponent_GracefulShutdown(t *testing.T) {
	defer goleak.VerifyNone(t, testcontainersIgnorePatterns...)

	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Create PostgreSQL testcontainer
	container, err := postgrescontainer.Run(ctx,
		"postgres:16-alpine",
		postgrescontainer.WithDatabase("djinn_test"),
		postgrescontainer.WithUsername("postgres"),
		postgrescontainer.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second),
		),
	)
	require.NoError(t, err, "Failed to start PostgreSQL container")

	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Logf("Failed to terminate container: %v", err)
		}
	})

	// Get connection string
	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err, "Failed to get connection string")

	// Parse connection config
	config, err := pgxpool.ParseConfig(connStr)
	require.NoError(t, err, "Failed to parse config")

	config.MaxConns = 5
	config.MinConns = 1

	// Create component
	component := NewComponent(config, logger)
	assert.Equal(t, "database", component.Name())
	assert.Empty(t, component.Dependencies())

	// Start component
	err = component.Start(ctx)
	require.NoError(t, err, "Failed to start component")

	// Verify health check works
	err = component.HealthCheck(ctx)
	require.NoError(t, err, "Health check failed")

	// Verify pool is available
	pool := component.Pool()
	assert.NotNil(t, pool, "Pool should not be nil")

	// Verify queries are available
	queries := component.Queries()
	assert.NotNil(t, queries, "Queries should not be nil")

	// Test graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err = component.Stop(shutdownCtx)
	assert.NoError(t, err, "Failed to stop component")

	// Verify component is properly shutdown
	pool = component.Pool()
	assert.Nil(t, pool, "Pool should be nil after shutdown")

	queries = component.Queries()
	assert.Nil(t, queries, "Queries should be nil after shutdown")
}

// TestDatabaseComponent_TransactionTracking tests transaction tracking during shutdown
func TestDatabaseComponent_TransactionTracking(t *testing.T) {
	defer goleak.VerifyNone(t, testcontainersIgnorePatterns...)

	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Create PostgreSQL testcontainer
	container, err := postgrescontainer.Run(ctx,
		"postgres:16-alpine",
		postgrescontainer.WithDatabase("djinn_test"),
		postgrescontainer.WithUsername("postgres"),
		postgrescontainer.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second),
		),
	)
	require.NoError(t, err, "Failed to start PostgreSQL container")

	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Logf("Failed to terminate container: %v", err)
		}
	})

	// Get connection string
	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err, "Failed to get connection string")

	// Parse connection config
	config, err := pgxpool.ParseConfig(connStr)
	require.NoError(t, err, "Failed to parse config")

	config.MaxConns = 10
	config.MinConns = 2

	// Create and start component
	component := NewComponent(config, logger)
	err = component.Start(ctx)
	require.NoError(t, err, "Failed to start component")

	// Start a transaction
	tx, err := component.BeginTx(ctx)
	require.NoError(t, err, "Failed to begin transaction")

	// Verify transaction is tracked
	assert.Equal(t, 1, component.GetActiveTransactionCount(), "Should have 1 active transaction")

	// Commit transaction
	err = tx.Commit(ctx)
	require.NoError(t, err, "Failed to commit transaction")

	// Verify transaction is no longer tracked
	assert.Equal(t, 0, component.GetActiveTransactionCount(), "Should have 0 active transactions after commit")

	// Test rollback tracking
	tx, err = component.BeginTx(ctx)
	require.NoError(t, err, "Failed to begin transaction")
	assert.Equal(t, 1, component.GetActiveTransactionCount(), "Should have 1 active transaction")

	err = tx.Rollback(ctx)
	require.NoError(t, err, "Failed to rollback transaction")
	assert.Equal(t, 0, component.GetActiveTransactionCount(), "Should have 0 active transactions after rollback")

	// Stop component
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err = component.Stop(shutdownCtx)
	assert.NoError(t, err, "Failed to stop component")
}

// TestDatabaseComponent_ShutdownWithActiveTransactions tests graceful shutdown with active transactions
func TestDatabaseComponent_ShutdownWithActiveTransactions(t *testing.T) {
	defer goleak.VerifyNone(t, testcontainersIgnorePatterns...)

	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Create PostgreSQL testcontainer
	container, err := postgrescontainer.Run(ctx,
		"postgres:16-alpine",
		postgrescontainer.WithDatabase("djinn_test"),
		postgrescontainer.WithUsername("postgres"),
		postgrescontainer.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second),
		),
	)
	require.NoError(t, err, "Failed to start PostgreSQL container")

	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Logf("Failed to terminate container: %v", err)
		}
	})

	// Get connection string
	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err, "Failed to get connection string")

	// Parse connection config
	config, err := pgxpool.ParseConfig(connStr)
	require.NoError(t, err, "Failed to parse config")

	config.MaxConns = 10
	config.MinConns = 2

	// Create and start component
	component := NewComponent(config, logger)
	err = component.Start(ctx)
	require.NoError(t, err, "Failed to start component")

	// Start multiple transactions
	var txs []interface{ pgx.Tx }
	for i := 0; i < 3; i++ {
		tx, err := component.BeginTx(ctx)
		require.NoError(t, err, "Failed to begin transaction %d", i)
		txs = append(txs, tx)
	}

	// Verify all transactions are tracked
	assert.Equal(t, 3, component.GetActiveTransactionCount(), "Should have 3 active transactions")

	// Start shutdown in background
	shutdownDone := make(chan error, 1)
	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 5*time.Second)
	defer shutdownCancel()

	go func() {
		shutdownDone <- component.Stop(shutdownCtx)
	}()

	// Let shutdown start
	time.Sleep(100 * time.Millisecond)

	// Verify new transactions are rejected during shutdown
	_, err = component.BeginTx(ctx)
	assert.Error(t, err, "Should reject new transactions during shutdown")
	assert.Contains(t, err.Error(), "database is shutting down", "Error should indicate shutdown")

	// Complete transactions
	for i, tx := range txs {
		if i%2 == 0 {
			err = tx.Commit(ctx)
			assert.NoError(t, err, "Failed to commit transaction %d", i)
		} else {
			err = tx.Rollback(ctx)
			assert.NoError(t, err, "Failed to rollback transaction %d", i)
		}
	}

	// Wait for shutdown to complete
	select {
	case err := <-shutdownDone:
		assert.NoError(t, err, "Shutdown should complete successfully")
	case <-time.After(10 * time.Second):
		t.Fatal("Shutdown timed out")
	}

	// Verify all transactions are cleaned up
	assert.Equal(t, 0, component.GetActiveTransactionCount(), "Should have 0 active transactions after shutdown")
}

// TestDatabaseComponent_ForceShutdown tests forced shutdown when transactions don't complete in time
func TestDatabaseComponent_ForceShutdown(t *testing.T) {
	defer goleak.VerifyNone(t, testcontainersIgnorePatterns...)

	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Create PostgreSQL testcontainer
	container, err := postgrescontainer.Run(ctx,
		"postgres:16-alpine",
		postgrescontainer.WithDatabase("djinn_test"),
		postgrescontainer.WithUsername("postgres"),
		postgrescontainer.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second),
		),
	)
	require.NoError(t, err, "Failed to start PostgreSQL container")

	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Logf("Failed to terminate container: %v", err)
		}
	})

	// Get connection string
	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err, "Failed to get connection string")

	// Parse connection config
	config, err := pgxpool.ParseConfig(connStr)
	require.NoError(t, err, "Failed to parse config")

	// Create and start component
	component := NewComponent(config, logger)
	err = component.Start(ctx)
	require.NoError(t, err, "Failed to start component")

	// Start a transaction that we won't complete
	_, err = component.BeginTx(ctx)
	require.NoError(t, err, "Failed to begin transaction")

	// Verify transaction is tracked
	assert.Equal(t, 1, component.GetActiveTransactionCount(), "Should have 1 active transaction")

	// Force shutdown with very short timeout
	shutdownCtx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

	start := time.Now()
	err = component.Stop(shutdownCtx)
	duration := time.Since(start)

	// Should complete quickly due to timeout (within 1 second)
	assert.True(t, duration < time.Second, "Force shutdown should complete quickly")
	assert.NoError(t, err, "Shutdown should complete even with active transactions")
}

// TestDatabaseComponent_ConcurrentTransactions tests concurrent transaction handling
func TestDatabaseComponent_ConcurrentTransactions(t *testing.T) {
	defer goleak.VerifyNone(t, testcontainersIgnorePatterns...)

	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Create PostgreSQL testcontainer
	container, err := postgrescontainer.Run(ctx,
		"postgres:16-alpine",
		postgrescontainer.WithDatabase("djinn_test"),
		postgrescontainer.WithUsername("postgres"),
		postgrescontainer.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second),
		),
	)
	require.NoError(t, err, "Failed to start PostgreSQL container")

	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Logf("Failed to terminate container: %v", err)
		}
	})

	// Get connection string
	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err, "Failed to get connection string")

	// Parse connection config
	config, err := pgxpool.ParseConfig(connStr)
	require.NoError(t, err, "Failed to parse config")

	config.MaxConns = 10
	config.MinConns = 2

	// Create and start component
	component := NewComponent(config, logger)
	err = component.Start(ctx)
	require.NoError(t, err, "Failed to start component")

	// Test concurrent transaction creation and completion
	const numGoroutines = 10
	const transactionsPerGoroutine = 5

	var wg sync.WaitGroup
	errChan := make(chan error, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < transactionsPerGoroutine; j++ {
				// Start transaction
				tx, err := component.BeginTx(ctx)
				if err != nil {
					errChan <- err
					return
				}

				// Do some work (simulate by sleeping)
				time.Sleep(10 * time.Millisecond)

				// Commit or rollback alternately
				if j%2 == 0 {
					err = tx.Commit(ctx)
				} else {
					err = tx.Rollback(ctx)
				}

				if err != nil {
					errChan <- err
					return
				}
			}
		}(i)
	}

	wg.Wait()
	close(errChan)

	// Check for errors
	for err := range errChan {
		assert.NoError(t, err, "Concurrent transaction failed")
	}

	// Verify all transactions are completed
	assert.Equal(t, 0, component.GetActiveTransactionCount(), "All transactions should be completed")

	// Stop component
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err = component.Stop(shutdownCtx)
	assert.NoError(t, err, "Failed to stop component")
}

// TestDatabaseComponent_SafetyFixes tests the safety fixes for nil pointer access
func TestDatabaseComponent_SafetyFixes(t *testing.T) {
	defer goleak.VerifyNone(t) // No testcontainers needed for this test

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	config := &pgxpool.Config{}

	// Create component but don't start it
	component := NewComponent(config, logger)

	// Pool should be nil before start
	pool := component.Pool()
	assert.Nil(t, pool, "Pool should be nil before start")

	// Queries should be nil before start
	queries := component.Queries()
	assert.Nil(t, queries, "Queries should be nil before start")

	// Health check should fail when pool is nil
	err := component.HealthCheck(context.Background())
	assert.Error(t, err, "Health check should fail when pool is nil")
	assert.Contains(t, err.Error(), "pool not initialized", "Error should indicate uninitialized pool")

	// Transaction creation should fail
	_, err = component.BeginTx(context.Background())
	assert.Error(t, err, "Transaction creation should fail when pool is nil")
}