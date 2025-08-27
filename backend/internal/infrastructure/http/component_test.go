package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestHTTPComponent_GracefulShutdown(t *testing.T) {
	defer goleak.VerifyNone(t)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	
	// Create simple handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Use a fixed port for testing (find a free one)
	addr := "localhost:18080" // Use a specific port for testing
	component := NewComponent(addr, handler, logger, nil)

	ctx := context.Background()

	// Start the component
	err := component.Start(ctx)
	require.NoError(t, err)

	// Make a test request to verify server is running
	client := &http.Client{Timeout: 1 * time.Second}
	resp, err := client.Get(fmt.Sprintf("http://%s", addr))
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Test graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = component.Stop(shutdownCtx)
	assert.NoError(t, err)

	// Verify server is no longer accepting connections
	_, err = client.Get(fmt.Sprintf("http://%s", addr))
	assert.Error(t, err) // Should fail because server is down

	// No goroutines should leak
}

func TestHTTPComponent_ShutdownWithActiveConnections(t *testing.T) {
	defer goleak.VerifyNone(t)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	
	// Handler that takes time to respond
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second) // Simulate slow response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Slow response"))
	})

	addr := "localhost:18081"
	component := NewComponent(addr, handler, logger, nil)

	ctx := context.Background()
	err := component.Start(ctx)
	require.NoError(t, err)

	// Start a slow request
	var wg sync.WaitGroup
	var requestErr error
	wg.Add(1)
	go func() {
		defer wg.Done()
		client := &http.Client{Timeout: 5 * time.Second}
		_, requestErr = client.Get(fmt.Sprintf("http://%s", addr))
	}()

	// Give request time to start
	time.Sleep(100 * time.Millisecond)

	// Start shutdown while request is active
	shutdownStart := time.Now()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = component.Stop(shutdownCtx)
	shutdownDuration := time.Since(shutdownStart)

	// Wait for the request to complete
	wg.Wait()

	// Shutdown should succeed
	assert.NoError(t, err)
	
	// Request should have completed successfully (graceful shutdown waits)
	assert.NoError(t, requestErr)
	
	// Shutdown should have waited for the request to complete
	assert.GreaterOrEqual(t, shutdownDuration, 2*time.Second)
	assert.Less(t, shutdownDuration, 5*time.Second) // But not hang forever
}

func TestHTTPComponent_ShutdownTimeout(t *testing.T) {
	defer goleak.VerifyNone(t)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	
	// Handler that takes very long time to respond
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(30 * time.Second) // Much longer than shutdown timeout
		w.WriteHeader(http.StatusOK)
	})

	addr := "localhost:18082"
	component := NewComponent(addr, handler, logger, nil)

	ctx := context.Background()
	err := component.Start(ctx)
	require.NoError(t, err)

	// Start a very slow request
	go func() {
		client := &http.Client{Timeout: 60 * time.Second}
		client.Get(fmt.Sprintf("http://%s", addr))
	}()

	// Give request time to start
	time.Sleep(100 * time.Millisecond)

	// Shutdown with short timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	shutdownStart := time.Now()
	err = component.Stop(shutdownCtx)
	shutdownDuration := time.Since(shutdownStart)

	// Should complete within timeout (force shutdown)
	assert.Error(t, err) // Should return timeout error
	assert.Less(t, shutdownDuration, 2*time.Second)
	
	// Server should still be stopped despite the error
	client := &http.Client{Timeout: 1 * time.Second}
	_, err = client.Get(fmt.Sprintf("http://%s", addr))
	assert.Error(t, err) // Should fail because server is down
}

func TestHTTPComponent_StartupVerification(t *testing.T) {
	defer goleak.VerifyNone(t)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	addr := "localhost:18083"
	component := NewComponent(addr, handler, logger, nil)

	startTime := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := component.Start(ctx)
	startDuration := time.Since(startTime)

	require.NoError(t, err)
	
	// Should start quickly (our verification is lightweight)
	assert.Less(t, startDuration, 2*time.Second)

	// Server should actually be listening
	assert.NotNil(t, component.Server())
	
	// Clean shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	
	err = component.Stop(shutdownCtx)
	assert.NoError(t, err)
}