package postgres

import (
	"context"
	"sync"
	"testing"
	"time"

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

// Comprehensive goleak ignore patterns for all network operations
var allNetworkIgnorePatterns = []goleak.Option{
	goleak.IgnoreTopFunction("github.com/testcontainers/testcontainers-go.(*Reaper).connect.func1"),
	goleak.IgnoreTopFunction("net.doBlockingWithCtx.func1"),
	goleak.IgnoreTopFunction("net._C2func_getaddrinfo"),
	goleak.IgnoreTopFunction("github.com/jackc/pgx/v5/pgxpool.(*Pool).createIdleResources"),
	goleak.IgnoreTopFunction("net.doBlockingWithCtx[...].func1"),
	goleak.IgnoreTopFunction("net.doBlockingWithCtx[...]"),
	goleak.IgnoreTopFunction("net.cgoLookupHostIP.func1"),
	goleak.IgnoreTopFunction("net.cgoLookupIP.func1"),
}

// TestDatabaseComponent_GracefulShutdown tests database component lifecycle
func TestDatabaseComponent_GracefulShutdown(t *testing.T) {
	defer goleak.VerifyNone(t, testcontainersIgnorePatterns...)

	ctx := context.Background()

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

	// Test connection lifecycle
	config := Config{
		URL:             connStr,
		MaxConnections:  10,
		MinConnections:  2,
		MaxConnLifetime: 5 * time.Minute,
	}

	db, err := NewConnection(config)
	require.NoError(t, err, "Failed to create database connection")

	// Verify connection works
	err = db.Health(ctx)
	require.NoError(t, err, "Database health check failed")

	// Test graceful shutdown
	db.Close()

	// Verify connection is closed
	err = db.Health(ctx)
	assert.Error(t, err, "Health check should fail after close")
}

// TestDatabaseComponent_ConnectionPoolCleanup tests connection pool cleanup
func TestDatabaseComponent_ConnectionPoolCleanup(t *testing.T) {
	defer goleak.VerifyNone(t, testcontainersIgnorePatterns...)

	ctx := context.Background()

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

	// Test multiple connection cycles
	for i := 0; i < 3; i++ {
		t.Run("connection_cycle", func(t *testing.T) {
			config := Config{
				URL:             connStr,
				MaxConnections:  5,
				MinConnections:  1,
				MaxConnLifetime: 1 * time.Minute,
			}

			db, err := NewConnection(config)
			require.NoError(t, err, "Failed to create database connection")

			// Verify connection works
			err = db.Health(ctx)
			require.NoError(t, err, "Database health check failed")

			// Get pool stats before operations
			pool := db.GetPool()
			require.NotNil(t, pool, "Pool should not be nil")

			// Perform some operations to create connections
			for j := 0; j < 3; j++ {
				err = db.Ping(ctx)
				require.NoError(t, err, "Ping should succeed")
			}

			// Close connection
			db.Close()

			// Verify pool is closed (subsequent operations should fail)
			err = db.Ping(ctx)
			assert.Error(t, err, "Ping should fail after close")
		})
	}
}

// TestDatabaseComponent_ConcurrentOperations tests concurrent operations during shutdown
func TestDatabaseComponent_ConcurrentOperations(t *testing.T) {
	defer goleak.VerifyNone(t, testcontainersIgnorePatterns...)

	ctx := context.Background()

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

	config := Config{
		URL:             connStr,
		MaxConnections:  10,
		MinConnections:  2,
		MaxConnLifetime: 5 * time.Minute,
	}

	db, err := NewConnection(config)
	require.NoError(t, err, "Failed to create database connection")

	// Start concurrent operations
	var wg sync.WaitGroup
	done := make(chan struct{})
	
	// Start background operations
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case <-done:
					return
				default:
					// Perform ping operation
					pingCtx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
					err := db.Ping(pingCtx)
					cancel()
					
					// Error is expected after close, but should not panic
					if err != nil {
						// Connection closed, exit gracefully
						return
					}
					time.Sleep(10 * time.Millisecond)
				}
			}
		}(i)
	}

	// Let operations run for a bit
	time.Sleep(100 * time.Millisecond)

	// Signal shutdown and close database
	close(done)
	db.Close()

	// Wait for all goroutines to complete
	wg.Wait()
}

// TestDatabaseComponent_PgBouncerConnection tests PgBouncer connection cleanup
func TestDatabaseComponent_PgBouncerConnection(t *testing.T) {
	defer goleak.VerifyNone(t, testcontainersIgnorePatterns...)

	ctx := context.Background()

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

	// Test PgBouncer connection (without PgBouncer URL, should use direct connection)
	config := Config{
		URL:             connStr,
		PgBouncerURL:    "", // Empty PgBouncer URL should use direct connection
		MaxConnections:  5,
		MinConnections:  1,
		MaxConnLifetime: 1 * time.Minute,
	}

	db, err := NewPgBouncerConnection(config)
	require.NoError(t, err, "Failed to create PgBouncer connection")

	// Verify connection works
	err = db.Health(ctx)
	require.NoError(t, err, "Database health check failed")

	// Test graceful shutdown
	db.Close()

	// Verify connection is closed
	err = db.Health(ctx)
	assert.Error(t, err, "Health check should fail after close")
}

// TestDatabaseComponent_MultipleStartStopCycles tests multiple start/stop cycles
func TestDatabaseComponent_MultipleStartStopCycles(t *testing.T) {
	defer goleak.VerifyNone(t, testcontainersIgnorePatterns...)

	ctx := context.Background()

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

	// Test multiple cycles
	for cycle := 0; cycle < 5; cycle++ {
		t.Run("cycle", func(t *testing.T) {
			config := Config{
				URL:             connStr,
				MaxConnections:  3,
				MinConnections:  1,
				MaxConnLifetime: 30 * time.Second,
			}

			// Create connection
			db, err := NewConnection(config)
			require.NoError(t, err, "Failed to create database connection in cycle %d", cycle)

			// Verify connection works
			err = db.Health(ctx)
			require.NoError(t, err, "Database health check failed in cycle %d", cycle)

			// Perform some operations
			for i := 0; i < 5; i++ {
				err = db.Ping(ctx)
				require.NoError(t, err, "Ping failed in cycle %d operation %d", cycle, i)
			}

			// Close connection
			db.Close()

			// Verify cleanup worked
			err = db.Health(ctx)
			assert.Error(t, err, "Health check should fail after close in cycle %d", cycle)
		})
	}
}

// TestDatabaseComponent_ConnectionTimeout tests connection timeout handling
func TestDatabaseComponent_ConnectionTimeout(t *testing.T) {
	defer goleak.VerifyNone(t, allNetworkIgnorePatterns...) // Ignore network operations

	// Test with invalid host (should timeout quickly)
	config := Config{
		URL:             "postgres://postgres:postgres@invalid-host:5432/test",
		MaxConnections:  5,
		MinConnections:  1,
		MaxConnLifetime: 1 * time.Minute,
	}

	// This should fail to create connection due to invalid host
	db, err := NewConnection(config)
	assert.Error(t, err, "Connection should fail with invalid host")
	assert.Nil(t, db, "DB should be nil on connection failure")
}

// TestDatabaseComponent_SetPool tests dynamic pool setting
func TestDatabaseComponent_SetPool(t *testing.T) {
	defer goleak.VerifyNone(t) // No testcontainers needed for this test

	// Create a DB instance without a pool
	db := &DB{}

	// Verify initial state
	assert.Nil(t, db.GetPool(), "Pool should be nil initially")
	assert.Nil(t, db.Queries, "Queries should be nil initially")

	// Set pool to nil (should handle gracefully)
	db.SetPool(nil)
	assert.Nil(t, db.GetPool(), "Pool should remain nil")
	assert.Nil(t, db.Queries, "Queries should remain nil")

	// Note: Testing with actual pool would require container setup
	// This tests the nil handling path of SetPool
}