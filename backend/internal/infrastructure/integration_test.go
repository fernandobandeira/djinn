package infrastructure

import (
	"context"
	"log/slog"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/database"
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
	goleak.IgnoreTopFunction("github.com/jackc/pgx/v5/pgxpool.(*Pool).backgroundHealthCheck"),
}

// TestDatabaseComponentIntegration tests the complete database component with real workloads
func TestDatabaseComponentIntegration(t *testing.T) {
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

	// Get connection string and configure database
	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err, "Failed to get connection string")

	dbConfig, err := pgxpool.ParseConfig(connStr)
	require.NoError(t, err, "Failed to parse config")
	dbConfig.MaxConns = 10
	dbConfig.MinConns = 2

	// Create database component
	dbComponent := database.NewComponent(dbConfig, logger)

	// Test application startup
	t.Run("component_startup", func(t *testing.T) {
		startCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		err = dbComponent.Start(startCtx)
		require.NoError(t, err, "Failed to start database component")

		// Verify database is healthy
		err = dbComponent.HealthCheck(ctx)
		assert.NoError(t, err, "Database health check failed")

		// Verify pool is available
		pool := dbComponent.Pool()
		assert.NotNil(t, pool, "Pool should not be nil")

		// Verify queries are available
		queries := dbComponent.Queries()
		assert.NotNil(t, queries, "Queries should not be nil")
	})

	// Test concurrent database operations (restart component for this test)
	t.Run("concurrent_operations", func(t *testing.T) {
		// Restart component for this test since previous test may have shut it down
		err = dbComponent.Start(ctx)
		require.NoError(t, err, "Failed to restart database component")

		const numWorkers = 10
		const operationsPerWorker = 20

		var wg sync.WaitGroup
		errChan := make(chan error, numWorkers)
		
		// Metrics tracking
		var totalTxCompleted int32
		var totalTxRolledback int32

		// Start database workers
		for i := 0; i < numWorkers; i++ {
			wg.Add(1)
			go func(workerID int) {
				defer wg.Done()

				for j := 0; j < operationsPerWorker; j++ {
					// Start a database transaction
					tx, err := dbComponent.BeginTx(ctx)
					if err != nil {
						errChan <- err
						return
					}

					// Simulate work with different patterns
					workDuration := time.Duration(10+workerID*2) * time.Millisecond
					time.Sleep(workDuration)

					// Commit or rollback based on pattern
					if (workerID+j)%3 == 0 {
						// Rollback every third operation
						if err := tx.Rollback(ctx); err != nil {
							errChan <- err
							return
						}
						totalTxRolledback++
					} else {
						// Commit most operations
						if err := tx.Commit(ctx); err != nil {
							errChan <- err
							return
						}
						totalTxCompleted++
					}
				}
			}(i)
		}

		// Monitor transaction count while workers run
		stopMonitor := make(chan struct{})
		monitorDone := make(chan struct{})
		go func() {
			defer close(monitorDone)
			maxActiveTx := 0
			
			for {
				select {
				case <-stopMonitor:
					return
				case <-time.After(100 * time.Millisecond):
					activeTx := dbComponent.GetActiveTransactionCount()
					if activeTx > maxActiveTx {
						maxActiveTx = activeTx
					}
				}
			}
		}()

		wg.Wait()
		close(errChan)
		close(stopMonitor)  // Signal monitor to stop
		<-monitorDone      // Wait for monitor to complete

		// Check for errors
		for err := range errChan {
			assert.NoError(t, err, "Concurrent database operation failed")
		}

		// Verify all transactions are completed
		assert.Equal(t, 0, dbComponent.GetActiveTransactionCount(), "All transactions should be completed")
		
		// Verify we processed expected number of operations
		totalExpected := int32(numWorkers * operationsPerWorker)
		assert.Equal(t, totalExpected, totalTxCompleted+totalTxRolledback, 
			"Should have processed all expected transactions")

		t.Logf("Completed %d commits, %d rollbacks", totalTxCompleted, totalTxRolledback)
	})

	// Test graceful shutdown with active operations
	t.Run("graceful_shutdown_with_load", func(t *testing.T) {
		// Restart component for this test 
		err = dbComponent.Start(ctx)
		require.NoError(t, err, "Failed to restart database component")
		const numConcurrentOps = 5
		
		// Start long-running operations in background
		var wg sync.WaitGroup
		stopOps := make(chan struct{})
		opsCompleted := make(chan int, numConcurrentOps)

		for i := 0; i < numConcurrentOps; i++ {
			wg.Add(1)
			go func(workerID int) {
				defer wg.Done()
				completed := 0
				
				for {
					select {
					case <-stopOps:
						opsCompleted <- completed
						return
					default:
						// Start transaction
						tx, err := dbComponent.BeginTx(ctx)
						if err != nil {
							// Expected during shutdown
							opsCompleted <- completed
							return
						}
						
						// Simulate work
						time.Sleep(50 * time.Millisecond)
						
						// Complete transaction
						if err := tx.Commit(ctx); err == nil {
							completed++
						}
					}
				}
			}(i)
		}

		// Let operations run for a bit
		time.Sleep(200 * time.Millisecond)
		
		// Verify operations are running
		activeBeforeShutdown := dbComponent.GetActiveTransactionCount()
		t.Logf("Active transactions before shutdown: %d", activeBeforeShutdown)

		// Start graceful shutdown
		shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		start := time.Now()
		err = dbComponent.Stop(shutdownCtx)
		shutdownDuration := time.Since(start)

		// Signal background operations to stop
		close(stopOps)
		wg.Wait()

		// Collect completion stats
		close(opsCompleted)
		totalOpsCompleted := 0
		for completed := range opsCompleted {
			totalOpsCompleted += completed
		}

		// Verify shutdown completed successfully
		assert.NoError(t, err, "Graceful shutdown should succeed")
		assert.True(t, shutdownDuration < 8*time.Second, "Shutdown should complete within reasonable time")

		// Verify database is shut down
		pool := dbComponent.Pool()
		assert.Nil(t, pool, "Database pool should be nil after shutdown")

		queries := dbComponent.Queries()
		assert.Nil(t, queries, "Queries should be nil after shutdown")

		// Health check should fail after shutdown
		err = dbComponent.HealthCheck(ctx)
		assert.Error(t, err, "Health check should fail after shutdown")

		t.Logf("Shutdown completed in %v, %d operations completed during test", 
			shutdownDuration, totalOpsCompleted)
	})

	// Test forced shutdown scenario
	t.Run("forced_shutdown", func(t *testing.T) {
		// Restart component for this test
		err = dbComponent.Start(ctx)
		require.NoError(t, err, "Failed to restart component")

		// Start transactions that we won't complete
		const numHangingTx = 3
		hangingTxs := make([]interface{}, numHangingTx)
		
		for i := 0; i < numHangingTx; i++ {
			tx, err := dbComponent.BeginTx(ctx)
			require.NoError(t, err, "Failed to start hanging transaction %d", i)
			hangingTxs[i] = tx
		}

		// Verify transactions are active
		assert.Equal(t, numHangingTx, dbComponent.GetActiveTransactionCount(), 
			"Should have hanging transactions")

		// Force shutdown with very short timeout
		shutdownCtx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
		defer cancel()

		start := time.Now()
		err = dbComponent.Stop(shutdownCtx)
		duration := time.Since(start)

		// Should complete quickly due to timeout
		assert.True(t, duration < 2*time.Second, "Force shutdown should complete quickly")
		assert.NoError(t, err, "Shutdown should complete even with hanging transactions")

		// Verify forced shutdown worked
		pool := dbComponent.Pool()
		assert.Nil(t, pool, "Pool should be nil after forced shutdown")

		t.Logf("Forced shutdown completed in %v", duration)

		// Clean up hanging transactions (they should fail now)
		for i, tx := range hangingTxs {
			if tx != nil {
				if err := tx.(interface{ Rollback(context.Context) error }).Rollback(ctx); err != nil {
					t.Logf("Expected error rolling back hanging tx %d: %v", i, err)
				}
			}
		}
	})
}

// TestDatabaseComponentErrorRecovery tests error scenarios and recovery
func TestDatabaseComponentErrorRecovery(t *testing.T) {
	defer goleak.VerifyNone(t, testcontainersIgnorePatterns...)

	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Test startup failure with invalid config
	t.Run("startup_failure", func(t *testing.T) {
		// Create database component with invalid config
		invalidConfig := &pgxpool.Config{}
		invalidConfig.ConnConfig.Host = "invalid-host"
		invalidConfig.ConnConfig.Port = 5432
		invalidConfig.ConnConfig.Database = "invalid"

		dbComponent := database.NewComponent(invalidConfig, logger)

		// Startup should fail
		startCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
		defer cancel()

		err := dbComponent.Start(startCtx)
		assert.Error(t, err, "Startup should fail with invalid database config")

		// Component should handle shutdown gracefully even after failed start
		shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 1*time.Second)
		defer shutdownCancel()
		
		err = dbComponent.Stop(shutdownCtx)
		assert.NoError(t, err, "Stop should succeed even after failed start")
	})

	// Test behavior with nil safety
	t.Run("nil_safety", func(t *testing.T) {
		config := &pgxpool.Config{}
		dbComponent := database.NewComponent(config, logger)

		// Component should handle nil pool gracefully
		pool := dbComponent.Pool()
		assert.Nil(t, pool, "Pool should be nil before start")

		queries := dbComponent.Queries()
		assert.Nil(t, queries, "Queries should be nil before start")

		// Health check should fail gracefully
		err := dbComponent.HealthCheck(ctx)
		assert.Error(t, err, "Health check should fail when pool is nil")

		// Transaction creation should fail gracefully
		_, err = dbComponent.BeginTx(ctx)
		assert.Error(t, err, "Transaction creation should fail when pool is nil")
	})
}

// TestDatabaseComponentStressTest performs stress testing of the database component
func TestDatabaseComponentStressTest(t *testing.T) {
	defer goleak.VerifyNone(t, testcontainersIgnorePatterns...)

	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

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

	// Configure high-performance database
	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err, "Failed to get connection string")

	dbConfig, err := pgxpool.ParseConfig(connStr)
	require.NoError(t, err, "Failed to parse config")
	dbConfig.MaxConns = 20
	dbConfig.MinConns = 5

	// Create and start component
	dbComponent := database.NewComponent(dbConfig, logger)
	err = dbComponent.Start(ctx)
	require.NoError(t, err, "Failed to start component")

	// Stress test with high concurrency
	t.Run("high_concurrency_stress", func(t *testing.T) {
		const numWorkers = 50
		const operationsPerWorker = 100
		const testDuration = 30 * time.Second

		var wg sync.WaitGroup
		startTime := time.Now()
		stopTime := startTime.Add(testDuration)
		
		errorCount := int64(0)
		successCount := int64(0)

		// Start stress workers
		for i := 0; i < numWorkers; i++ {
			wg.Add(1)
			go func(workerID int) {
				defer wg.Done()
				
				localErrors := 0
				localSuccesses := 0
				
				for j := 0; j < operationsPerWorker && time.Now().Before(stopTime); j++ {
					// Random operation pattern
					tx, err := dbComponent.BeginTx(ctx)
					if err != nil {
						localErrors++
						continue
					}

					// Variable work duration
					workTime := time.Duration(5+j%20) * time.Millisecond
					time.Sleep(workTime)

					// Random commit/rollback pattern
					if j%7 == 0 {
						err = tx.Rollback(ctx)
					} else {
						err = tx.Commit(ctx)
					}

					if err != nil {
						localErrors++
					} else {
						localSuccesses++
					}
				}
				
				t.Logf("Worker %d completed: %d successes, %d errors", 
					workerID, localSuccesses, localErrors)
				
				successCount += int64(localSuccesses)
				errorCount += int64(localErrors)
			}(i)
		}

		wg.Wait()

		// Verify stress test results
		totalOps := successCount + errorCount
		assert.Greater(t, totalOps, int64(1000), "Should have completed significant number of operations")
		assert.Equal(t, 0, dbComponent.GetActiveTransactionCount(), "All transactions should be completed")
		
		successRate := float64(successCount) / float64(totalOps) * 100
		assert.Greater(t, successRate, 95.0, "Success rate should be high")

		t.Logf("Stress test completed: %d total operations, %.2f%% success rate", 
			totalOps, successRate)
	})

	// Test graceful shutdown under stress
	shutdownCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	
	err = dbComponent.Stop(shutdownCtx)
	assert.NoError(t, err, "Shutdown should succeed even after stress test")
}