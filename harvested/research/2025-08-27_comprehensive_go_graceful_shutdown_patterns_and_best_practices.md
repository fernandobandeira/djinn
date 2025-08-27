# Go Graceful Shutdown Patterns and Best Practices

**Research Date**: 2025-08-27  
**Research Type**: Comprehensive Technical Documentation  
**Source Agent**: architect  
**Collection**: go-patterns  

## Executive Summary

This document provides comprehensive guidance for implementing graceful shutdown patterns in Go applications, covering industry-standard libraries, best practices for HTTP servers, database connection management, background goroutine lifecycle management, and integration with container orchestration platforms.

## Table of Contents

1. [Core Concepts](#core-concepts)
2. [Signal Handling Patterns](#signal-handling-patterns)
3. [HTTP Server Graceful Shutdown](#http-server-graceful-shutdown)
4. [Database Connection Draining](#database-connection-draining)
5. [Background Goroutine Management](#background-goroutine-management)
6. [Industry-Standard Libraries](#industry-standard-libraries)
7. [Container Orchestration Integration](#container-orchestration-integration)
8. [Common Pitfalls and Race Conditions](#common-pitfalls-and-race-conditions)
9. [Monitoring Integration](#monitoring-integration)
10. [Framework-Specific Patterns](#framework-specific-patterns)
11. [Complete Implementation Examples](#complete-implementation-examples)

## Core Concepts

### Understanding Graceful Shutdown

Graceful shutdown is the process of cleanly terminating a Go application by:
- Stopping the acceptance of new requests/connections
- Allowing in-flight requests to complete within a reasonable timeout
- Properly closing database connections and other resources
- Cleaning up background goroutines and workers
- Saving application state if necessary
- Exiting with appropriate status codes

### Signal Types

```go
// Primary signals for graceful shutdown
SIGTERM  // Termination signal (default for docker stop)
SIGINT   // Interrupt signal (Ctrl+C)
SIGUSR1  // User-defined signal 1 (custom operations)
SIGUSR2  // User-defined signal 2 (custom operations)

// Immediate termination (cannot be caught)
SIGKILL  // Force kill (docker stop -f)
```

## Signal Handling Patterns

### Basic Signal Handling

```go
package main

import (
    "context"
    "log/slog"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    // Create context that can be canceled
    ctx, cancel := context.WithCancel(context.Background())
    
    // Channel to listen for interrupt signals
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    
    // Start application components
    go startHTTPServer(ctx)
    go startBackgroundWorkers(ctx)
    
    // Wait for signal
    sig := <-sigChan
    slog.Info("Received signal, starting graceful shutdown", "signal", sig)
    
    // Cancel context to signal all components to stop
    cancel()
    
    // Give components time to shut down gracefully
    shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer shutdownCancel()
    
    // Perform graceful shutdown operations
    if err := performGracefulShutdown(shutdownCtx); err != nil {
        slog.Error("Error during graceful shutdown", "error", err)
        os.Exit(1)
    }
    
    slog.Info("Application shut down gracefully")
}
```

### Advanced Signal Handling with Multiple Signals

```go
func setupSignalHandling() <-chan struct{} {
    signals := make(chan os.Signal, 1)
    done := make(chan struct{})
    
    // Handle multiple signal types
    signal.Notify(signals, 
        syscall.SIGINT,  // Ctrl+C
        syscall.SIGTERM, // Docker stop
        syscall.SIGUSR1, // Custom reload
        syscall.SIGUSR2, // Custom debug
    )
    
    go func() {
        defer close(done)
        
        for {
            sig := <-signals
            switch sig {
            case syscall.SIGUSR1:
                slog.Info("Reloading configuration")
                handleConfigReload()
                continue
            case syscall.SIGUSR2:
                slog.Info("Dumping debug information")
                handleDebugDump()
                continue
            case syscall.SIGINT, syscall.SIGTERM:
                slog.Info("Received shutdown signal", "signal", sig)
                return
            }
        }
    }()
    
    return done
}
```

## HTTP Server Graceful Shutdown

### Standard HTTP Server Shutdown

```go
package main

import (
    "context"
    "log/slog"
    "net/http"
    "time"
)

func startHTTPServer(ctx context.Context) error {
    server := &http.Server{
        Addr:         ":8080",
        Handler:      setupRoutes(),
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  60 * time.Second,
    }
    
    // Start server in goroutine
    go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            slog.Error("HTTP server error", "error", err)
        }
    }()
    
    slog.Info("HTTP server started", "addr", server.Addr)
    
    // Wait for context cancellation
    <-ctx.Done()
    
    // Create shutdown context with timeout
    shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()
    
    slog.Info("Shutting down HTTP server...")
    
    // Graceful shutdown
    if err := server.Shutdown(shutdownCtx); err != nil {
        slog.Error("HTTP server shutdown error", "error", err)
        return err
    }
    
    slog.Info("HTTP server shut down successfully")
    return nil
}
```

### Chi Framework Graceful Shutdown

```go
package main

import (
    "context"
    "net/http"
    "time"
    
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func startChiServer(ctx context.Context) error {
    r := chi.NewRouter()
    
    // Middleware stack
    r.Use(middleware.RequestID)
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(middleware.Timeout(30 * time.Second))
    
    // Health check endpoint
    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })
    
    // Graceful shutdown endpoint for testing
    r.Post("/shutdown", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusAccepted)
        w.Write([]byte("Shutdown initiated"))
        go func() {
            time.Sleep(1 * time.Second)
            // Trigger shutdown (in real app, this would be through signal)
        }()
    })
    
    server := &http.Server{
        Addr:    ":8080",
        Handler: r,
    }
    
    // Server goroutine
    serverErr := make(chan error, 1)
    go func() {
        serverErr <- server.ListenAndServe()
    }()
    
    // Wait for context done or server error
    select {
    case err := <-serverErr:
        if err != http.ErrServerClosed {
            return err
        }
    case <-ctx.Done():
        // Graceful shutdown
        shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
        defer cancel()
        
        return server.Shutdown(shutdownCtx)
    }
    
    return nil
}
```

## Database Connection Draining

### PostgreSQL Connection Draining (pgx)

```go
package main

import (
    "context"
    "log/slog"
    "time"
    
    "github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseManager struct {
    pool *pgxpool.Pool
}

func NewDatabaseManager(databaseURL string) (*DatabaseManager, error) {
    config, err := pgxpool.ParseConfig(databaseURL)
    if err != nil {
        return nil, err
    }
    
    // Configure connection pool for graceful shutdown
    config.MaxConns = 25
    config.MinConns = 5
    config.MaxConnLifetime = time.Hour
    config.MaxConnIdleTime = time.Minute * 30
    
    pool, err := pgxpool.NewWithConfig(context.Background(), config)
    if err != nil {
        return nil, err
    }
    
    return &DatabaseManager{pool: pool}, nil
}

func (dm *DatabaseManager) GracefulShutdown(ctx context.Context) error {
    slog.Info("Starting database connection draining...")
    
    // Stop accepting new connections
    dm.pool.Config().MaxConns = 0
    
    // Wait for active connections to finish or timeout
    done := make(chan struct{})
    go func() {
        defer close(done)
        
        // Monitor active connections
        ticker := time.NewTicker(1 * time.Second)
        defer ticker.Stop()
        
        for {
            select {
            case <-ctx.Done():
                return
            case <-ticker.C:
                stats := dm.pool.Stat()
                if stats.AcquiredConns() == 0 {
                    slog.Info("All database connections closed")
                    return
                }
                slog.Info("Waiting for database connections", 
                    "active", stats.AcquiredConns(),
                    "idle", stats.IdleConns(),
                    "total", stats.TotalConns(),
                )
            }
        }
    }()
    
    // Wait for connections to drain or timeout
    select {
    case <-done:
        slog.Info("Database connections drained successfully")
    case <-ctx.Done():
        slog.Warn("Database connection draining timed out")
    }
    
    // Force close remaining connections
    dm.pool.Close()
    slog.Info("Database pool closed")
    
    return nil
}

// Health check that respects shutdown state
func (dm *DatabaseManager) HealthCheck(ctx context.Context) error {
    if dm.pool == nil {
        return errors.New("database pool is nil")
    }
    
    // Quick ping with timeout
    pingCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
    defer cancel()
    
    return dm.pool.Ping(pingCtx)
}
```

### Database Transaction Cleanup

```go
func (dm *DatabaseManager) WithTransaction(ctx context.Context, fn func(tx pgx.Tx) error) error {
    tx, err := dm.pool.Begin(ctx)
    if err != nil {
        return err
    }
    
    defer func() {
        if p := recover(); p != nil {
            _ = tx.Rollback(ctx)
            panic(p)
        } else if err != nil {
            _ = tx.Rollback(ctx)
        } else {
            err = tx.Commit(ctx)
        }
    }()
    
    // Check for context cancellation before proceeding
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
    }
    
    err = fn(tx)
    return err
}
```

## Background Goroutine Management

### Worker Pool Pattern

```go
package main

import (
    "context"
    "log/slog"
    "sync"
    "time"
)

type WorkerPool struct {
    workers    int
    jobQueue   chan Job
    wg         sync.WaitGroup
    ctx        context.Context
    cancel     context.CancelFunc
}

type Job interface {
    Process(ctx context.Context) error
}

func NewWorkerPool(workers int, queueSize int) *WorkerPool {
    ctx, cancel := context.WithCancel(context.Background())
    
    return &WorkerPool{
        workers:  workers,
        jobQueue: make(chan Job, queueSize),
        ctx:      ctx,
        cancel:   cancel,
    }
}

func (wp *WorkerPool) Start() {
    slog.Info("Starting worker pool", "workers", wp.workers)
    
    for i := 0; i < wp.workers; i++ {
        wp.wg.Add(1)
        go wp.worker(i)
    }
}

func (wp *WorkerPool) worker(id int) {
    defer wp.wg.Done()
    
    slog.Info("Worker started", "worker_id", id)
    
    for {
        select {
        case <-wp.ctx.Done():
            slog.Info("Worker stopping", "worker_id", id)
            return
        case job, ok := <-wp.jobQueue:
            if !ok {
                slog.Info("Job queue closed, worker stopping", "worker_id", id)
                return
            }
            
            // Process job with timeout
            jobCtx, cancel := context.WithTimeout(wp.ctx, 30*time.Second)
            if err := job.Process(jobCtx); err != nil {
                slog.Error("Job processing failed", 
                    "worker_id", id, 
                    "error", err)
            }
            cancel()
        }
    }
}

func (wp *WorkerPool) Submit(job Job) bool {
    select {
    case wp.jobQueue <- job:
        return true
    case <-wp.ctx.Done():
        return false
    default:
        // Queue full, job rejected
        return false
    }
}

func (wp *WorkerPool) GracefulShutdown(timeout time.Duration) error {
    slog.Info("Starting worker pool graceful shutdown")
    
    // Stop accepting new jobs
    wp.cancel()
    
    // Close job queue to signal workers
    close(wp.jobQueue)
    
    // Wait for workers to finish with timeout
    done := make(chan struct{})
    go func() {
        defer close(done)
        wp.wg.Wait()
    }()
    
    select {
    case <-done:
        slog.Info("Worker pool shut down gracefully")
        return nil
    case <-time.After(timeout):
        slog.Warn("Worker pool shutdown timed out")
        return errors.New("worker pool shutdown timeout")
    }
}
```

### Background Service Pattern

```go
type BackgroundService struct {
    name     string
    interval time.Duration
    task     func(ctx context.Context) error
    ctx      context.Context
    cancel   context.CancelFunc
    wg       sync.WaitGroup
}

func NewBackgroundService(name string, interval time.Duration, task func(ctx context.Context) error) *BackgroundService {
    ctx, cancel := context.WithCancel(context.Background())
    
    return &BackgroundService{
        name:     name,
        interval: interval,
        task:     task,
        ctx:      ctx,
        cancel:   cancel,
    }
}

func (bs *BackgroundService) Start() {
    bs.wg.Add(1)
    go bs.run()
}

func (bs *BackgroundService) run() {
    defer bs.wg.Done()
    
    slog.Info("Background service started", "service", bs.name)
    
    ticker := time.NewTicker(bs.interval)
    defer ticker.Stop()
    
    // Run immediately on start
    if err := bs.task(bs.ctx); err != nil {
        slog.Error("Background task failed", 
            "service", bs.name, 
            "error", err)
    }
    
    for {
        select {
        case <-bs.ctx.Done():
            slog.Info("Background service stopping", "service", bs.name)
            return
        case <-ticker.C:
            if err := bs.task(bs.ctx); err != nil {
                slog.Error("Background task failed", 
                    "service", bs.name, 
                    "error", err)
            }
        }
    }
}

func (bs *BackgroundService) GracefulShutdown(timeout time.Duration) error {
    slog.Info("Shutting down background service", "service", bs.name)
    
    bs.cancel()
    
    done := make(chan struct{})
    go func() {
        defer close(done)
        bs.wg.Wait()
    }()
    
    select {
    case <-done:
        slog.Info("Background service shut down", "service", bs.name)
        return nil
    case <-time.After(timeout):
        slog.Warn("Background service shutdown timed out", "service", bs.name)
        return errors.New("background service shutdown timeout")
    }
}
```

## Industry-Standard Libraries

### 1. github.com/oklog/run

Best for coordinating multiple components with automatic cleanup.

```go
package main

import (
    "context"
    "log/slog"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    
    "github.com/oklog/run"
)

func main() {
    var g run.Group
    
    // HTTP server component
    {
        server := &http.Server{Addr: ":8080"}
        g.Add(func() error {
            return server.ListenAndServe()
        }, func(error) {
            ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
            defer cancel()
            server.Shutdown(ctx)
        })
    }
    
    // Background worker component  
    {
        ctx, cancel := context.WithCancel(context.Background())
        g.Add(func() error {
            return runBackgroundWorker(ctx)
        }, func(error) {
            cancel()
        })
    }
    
    // Signal handling component
    {
        sigChan := make(chan os.Signal, 1)
        signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
        g.Add(func() error {
            sig := <-sigChan
            slog.Info("Received signal", "signal", sig)
            return nil
        }, func(error) {
            close(sigChan)
        })
    }
    
    slog.Info("Starting application...")
    if err := g.Run(); err != nil {
        slog.Error("Application error", "error", err)
        os.Exit(1)
    }
    
    slog.Info("Application stopped")
}
```

### 2. github.com/vardius/shutdown

Provides a more structured approach to graceful shutdown.

```go
package main

import (
    "context"
    "log/slog"
    "net/http"
    "time"
    
    "github.com/vardius/shutdown"
)

func main() {
    ctx, stop := shutdown.WithInterrupt(context.Background())
    defer stop()
    
    // HTTP Server
    server := &http.Server{
        Addr:    ":8080",
        Handler: setupRoutes(),
    }
    
    go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            slog.Error("HTTP server error", "error", err)
        }
    }()
    
    // Background services
    workerPool := NewWorkerPool(5, 100)
    workerPool.Start()
    
    // Wait for shutdown signal
    <-ctx.Done()
    
    slog.Info("Starting graceful shutdown...")
    
    // Shutdown with timeout
    shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    // Shutdown HTTP server
    if err := server.Shutdown(shutdownCtx); err != nil {
        slog.Error("HTTP server shutdown error", "error", err)
    }
    
    // Shutdown worker pool
    if err := workerPool.GracefulShutdown(20 * time.Second); err != nil {
        slog.Error("Worker pool shutdown error", "error", err)
    }
    
    slog.Info("Graceful shutdown completed")
}
```

### 3. github.com/uber-go/fx (Dependency Injection with Lifecycle)

Best for complex applications with many dependencies.

```go
package main

import (
    "context"
    "log/slog"
    "net/http"
    
    "go.uber.org/fx"
)

func NewHTTPServer(lc fx.Lifecycle) *http.Server {
    server := &http.Server{
        Addr:    ":8080",
        Handler: setupRoutes(),
    }
    
    lc.Append(fx.Hook{
        OnStart: func(ctx context.Context) error {
            go func() {
                if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
                    slog.Error("HTTP server error", "error", err)
                }
            }()
            slog.Info("HTTP server started")
            return nil
        },
        OnStop: func(ctx context.Context) error {
            slog.Info("Shutting down HTTP server...")
            return server.Shutdown(ctx)
        },
    })
    
    return server
}

func NewWorkerPool(lc fx.Lifecycle) *WorkerPool {
    pool := NewWorkerPool(5, 100)
    
    lc.Append(fx.Hook{
        OnStart: func(ctx context.Context) error {
            pool.Start()
            slog.Info("Worker pool started")
            return nil
        },
        OnStop: func(ctx context.Context) error {
            slog.Info("Shutting down worker pool...")
            return pool.GracefulShutdown(15 * time.Second)
        },
    })
    
    return pool
}

func main() {
    fx.New(
        fx.Provide(
            NewHTTPServer,
            NewWorkerPool,
        ),
        fx.Invoke(func(*http.Server, *WorkerPool) {
            // Dependencies will be started automatically
        }),
    ).Run()
}
```

## Container Orchestration Integration

### Kubernetes Integration

```yaml
# deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-app
spec:
  template:
    spec:
      containers:
      - name: app
        image: go-app:latest
        lifecycle:
          preStop:
            exec:
              command: ["/bin/sh", "-c", "sleep 10"]
        # Kubernetes sends SIGTERM, then waits for terminationGracePeriodSeconds
        terminationGracePeriodSeconds: 45
        readinessProbe:
          httpGet:
            path: /health
            port: 8080
          periodSeconds: 10
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          periodSeconds: 30
          failureThreshold: 3
```

```go
// Kubernetes-aware health check
func kubernetesHealthCheck(w http.ResponseWriter, r *http.Request) {
    // Check if application is shutting down
    select {
    case <-shutdownCtx.Done():
        // Return 503 to signal readiness probe failure
        w.WriteHeader(http.StatusServiceUnavailable)
        w.Write([]byte("Shutting down"))
        return
    default:
    }
    
    // Perform actual health checks
    if err := performHealthChecks(); err != nil {
        w.WriteHeader(http.StatusServiceUnavailable)
        w.Write([]byte(err.Error()))
        return
    }
    
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}
```

### Docker Integration

```dockerfile
# Dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/

COPY --from=builder /app/app .

# Use tini as init system to properly handle signals
RUN apk add --no-cache tini
ENTRYPOINT ["/sbin/tini", "--"]

EXPOSE 8080
CMD ["./app"]
```

```bash
# Docker Compose with proper signal handling
version: '3.8'
services:
  app:
    build: .
    ports:
      - "8080:8080"
    stop_signal: SIGTERM
    stop_grace_period: 45s
    healthcheck:
      test: ["CMD", "wget", "--no-verbose", "--tries=1", "--spider", "http://localhost:8080/health"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s
```

## Common Pitfalls and Race Conditions

### 1. Race Condition in Signal Handling

**Problem**: Multiple goroutines trying to handle the same signal.

```go
// INCORRECT - Race condition
func badSignalHandling() {
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGTERM)
    
    // Multiple goroutines reading from the same channel
    go func() {
        <-sigChan
        // First handler
    }()
    
    go func() {
        <-sigChan  // This might never receive if first handler got it
        // Second handler
    }()
}

// CORRECT - Single signal handler with coordination
func goodSignalHandling() {
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGTERM)
    
    ctx, cancel := context.WithCancel(context.Background())
    
    go func() {
        <-sigChan
        cancel() // Signal all components via context
    }()
    
    // All components watch the context
    go componentA(ctx)
    go componentB(ctx)
}
```

### 2. Deadlock During Shutdown

**Problem**: Waiting for resources that are blocked.

```go
// INCORRECT - Can deadlock
func badShutdown() {
    var wg sync.WaitGroup
    
    wg.Add(1)
    go func() {
        defer wg.Done()
        // This might block forever
        for {
            select {
            case data := <-someChan:
                process(data)
            }
        }
    }()
    
    // This will wait forever if someChan never closes
    wg.Wait()
}

// CORRECT - Use context for cancellation
func goodShutdown() {
    var wg sync.WaitGroup
    ctx, cancel := context.WithCancel(context.Background())
    
    wg.Add(1)
    go func() {
        defer wg.Done()
        for {
            select {
            case <-ctx.Done():
                return // Exit on cancellation
            case data := <-someChan:
                process(data)
            }
        }
    }()
    
    // Trigger shutdown
    cancel()
    
    // Wait with timeout
    done := make(chan struct{})
    go func() {
        wg.Wait()
        close(done)
    }()
    
    select {
    case <-done:
        // Clean shutdown
    case <-time.After(30 * time.Second):
        // Force exit after timeout
    }
}
```

### 3. Resource Leaks

**Problem**: Not properly closing resources during shutdown.

```go
// Resource manager with proper cleanup
type ResourceManager struct {
    resources []io.Closer
    mu        sync.RWMutex
}

func (rm *ResourceManager) Add(r io.Closer) {
    rm.mu.Lock()
    defer rm.mu.Unlock()
    rm.resources = append(rm.resources, r)
}

func (rm *ResourceManager) Close() error {
    rm.mu.Lock()
    defer rm.mu.Unlock()
    
    var errs []error
    for i := len(rm.resources) - 1; i >= 0; i-- {
        if err := rm.resources[i].Close(); err != nil {
            errs = append(errs, err)
        }
    }
    
    rm.resources = rm.resources[:0] // Clear slice
    
    if len(errs) > 0 {
        return fmt.Errorf("resource cleanup errors: %v", errs)
    }
    return nil
}
```

## Monitoring Integration

### OpenTelemetry Shutdown Monitoring

```go
package main

import (
    "context"
    "log/slog"
    
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/metric"
)

type ShutdownMonitor struct {
    shutdownDuration metric.Float64Histogram
    shutdownCounter  metric.Int64Counter
    shutdownGauge    metric.Int64UpDownCounter
}

func NewShutdownMonitor() (*ShutdownMonitor, error) {
    meter := otel.Meter("app/shutdown")
    
    shutdownDuration, err := meter.Float64Histogram(
        "app_shutdown_duration_seconds",
        metric.WithDescription("Time taken for graceful shutdown"),
    )
    if err != nil {
        return nil, err
    }
    
    shutdownCounter, err := meter.Int64Counter(
        "app_shutdown_total",
        metric.WithDescription("Total number of shutdown attempts"),
    )
    if err != nil {
        return nil, err
    }
    
    shutdownGauge, err := meter.Int64UpDownCounter(
        "app_shutdown_in_progress",
        metric.WithDescription("Whether shutdown is in progress"),
    )
    if err != nil {
        return nil, err
    }
    
    return &ShutdownMonitor{
        shutdownDuration: shutdownDuration,
        shutdownCounter:  shutdownCounter,
        shutdownGauge:    shutdownGauge,
    }, nil
}

func (sm *ShutdownMonitor) RecordShutdown(ctx context.Context, duration time.Duration, success bool) {
    attrs := []attribute.KeyValue{
        attribute.Bool("success", success),
    }
    
    sm.shutdownDuration.Record(ctx, duration.Seconds(), metric.WithAttributes(attrs...))
    sm.shutdownCounter.Add(ctx, 1, metric.WithAttributes(attrs...))
}

func (sm *ShutdownMonitor) SetShutdownInProgress(ctx context.Context, inProgress bool) {
    value := int64(0)
    if inProgress {
        value = 1
    }
    sm.shutdownGauge.Add(ctx, value, metric.WithAttributes())
}

// Usage in main function
func monitoredGracefulShutdown(ctx context.Context, monitor *ShutdownMonitor) error {
    start := time.Now()
    monitor.SetShutdownInProgress(ctx, true)
    defer func() {
        duration := time.Since(start)
        success := true // Determine based on error state
        monitor.RecordShutdown(ctx, duration, success)
        monitor.SetShutdownInProgress(ctx, false)
    }()
    
    // Perform shutdown operations
    return performShutdown(ctx)
}
```

### Prometheus Metrics

```go
package main

import (
    "time"
    
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
)

var (
    shutdownDuration = promauto.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "app_graceful_shutdown_duration_seconds",
            Help: "Time taken for graceful shutdown",
            Buckets: prometheus.DefBuckets,
        },
        []string{"component", "success"},
    )
    
    shutdownInProgress = promauto.NewGauge(
        prometheus.GaugeOpts{
            Name: "app_shutdown_in_progress",
            Help: "Whether graceful shutdown is in progress",
        },
    )
    
    activeConnections = promauto.NewGauge(
        prometheus.GaugeOpts{
            Name: "app_active_connections",
            Help: "Number of active connections during shutdown",
        },
    )
)

func instrumentedShutdown(component string, shutdownFunc func() error) error {
    start := time.Now()
    shutdownInProgress.Set(1)
    defer shutdownInProgress.Set(0)
    
    err := shutdownFunc()
    
    duration := time.Since(start).Seconds()
    success := "true"
    if err != nil {
        success = "false"
    }
    
    shutdownDuration.WithLabelValues(component, success).Observe(duration)
    return err
}
```

## Framework-Specific Patterns

### Gin Framework

```go
package main

import (
    "context"
    "log/slog"
    "net/http"
    "time"
    
    "github.com/gin-gonic/gin"
)

func startGinServer(ctx context.Context) error {
    gin.SetMode(gin.ReleaseMode)
    r := gin.New()
    
    // Middleware
    r.Use(gin.Recovery())
    r.Use(gin.Logger())
    
    // Routes
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })
    
    server := &http.Server{
        Addr:    ":8080",
        Handler: r,
    }
    
    // Start server
    go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            slog.Error("Gin server error", "error", err)
        }
    }()
    
    // Wait for context cancellation
    <-ctx.Done()
    
    // Graceful shutdown
    shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()
    
    return server.Shutdown(shutdownCtx)
}
```

### Echo Framework

```go
package main

import (
    "context"
    "log/slog"
    "time"
    
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func startEchoServer(ctx context.Context) error {
    e := echo.New()
    e.HideBanner = true
    
    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    
    // Routes
    e.GET("/health", func(c echo.Context) error {
        return c.JSON(200, map[string]string{"status": "ok"})
    })
    
    // Start server in goroutine
    go func() {
        if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
            slog.Error("Echo server error", "error", err)
        }
    }()
    
    // Wait for context cancellation
    <-ctx.Done()
    
    // Graceful shutdown
    shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()
    
    return e.Shutdown(shutdownCtx)
}
```

### Fiber Framework

```go
package main

import (
    "context"
    "log/slog"
    "time"
    
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/recover"
)

func startFiberServer(ctx context.Context) error {
    app := fiber.New(fiber.Config{
        DisableStartupMessage: true,
        GracefulShutdown:      15 * time.Second,
    })
    
    // Middleware
    app.Use(logger.New())
    app.Use(recover.New())
    
    // Routes
    app.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"status": "ok"})
    })
    
    // Start server in goroutine
    go func() {
        if err := app.Listen(":8080"); err != nil {
            slog.Error("Fiber server error", "error", err)
        }
    }()
    
    // Wait for context cancellation
    <-ctx.Done()
    
    // Fiber has built-in graceful shutdown
    return app.Shutdown()
}
```

## Complete Implementation Examples

### Production-Ready Application

```go
package main

import (
    "context"
    "errors"
    "log/slog"
    "net/http"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"
    
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
    server      *http.Server
    db          *pgxpool.Pool
    workerPool  *WorkerPool
    bgServices  []*BackgroundService
    monitor     *ShutdownMonitor
    
    shutdownOnce sync.Once
    shutdownChan chan struct{}
}

func NewApplication() *Application {
    return &Application{
        shutdownChan: make(chan struct{}),
    }
}

func (app *Application) Initialize() error {
    var err error
    
    // Initialize database
    app.db, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
    if err != nil {
        return err
    }
    
    // Initialize worker pool
    app.workerPool = NewWorkerPool(10, 1000)
    
    // Initialize HTTP server
    router := chi.NewRouter()
    router.Use(middleware.RequestID)
    router.Use(middleware.Logger)
    router.Use(middleware.Recoverer)
    router.Use(middleware.Timeout(30 * time.Second))
    
    // Health check with shutdown awareness
    router.Get("/health", app.healthCheck)
    router.Get("/ready", app.readinessCheck)
    
    app.server = &http.Server{
        Addr:    ":8080",
        Handler: router,
    }
    
    // Initialize background services
    app.bgServices = append(app.bgServices, 
        NewBackgroundService("cleanup", 5*time.Minute, app.cleanupTask),
        NewBackgroundService("metrics", 1*time.Minute, app.metricsTask),
    )
    
    // Initialize monitoring
    app.monitor, err = NewShutdownMonitor()
    if err != nil {
        return err
    }
    
    return nil
}

func (app *Application) Start() error {
    // Start worker pool
    app.workerPool.Start()
    slog.Info("Worker pool started")
    
    // Start background services
    for _, service := range app.bgServices {
        service.Start()
    }
    slog.Info("Background services started", "count", len(app.bgServices))
    
    // Start HTTP server
    go func() {
        slog.Info("HTTP server starting", "addr", app.server.Addr)
        if err := app.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            slog.Error("HTTP server error", "error", err)
        }
    }()
    
    return nil
}

func (app *Application) WaitForShutdown() {
    // Set up signal handling
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    
    select {
    case sig := <-sigChan:
        slog.Info("Received shutdown signal", "signal", sig)
    case <-app.shutdownChan:
        slog.Info("Received shutdown request")
    }
    
    app.Shutdown()
}

func (app *Application) Shutdown() {
    app.shutdownOnce.Do(func() {
        start := time.Now()
        app.monitor.SetShutdownInProgress(context.Background(), true)
        
        slog.Info("Starting graceful shutdown...")
        
        // Create shutdown context with timeout
        ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
        defer cancel()
        
        var wg sync.WaitGroup
        
        // Shutdown HTTP server
        wg.Add(1)
        go func() {
            defer wg.Done()
            if err := app.shutdownHTTPServer(ctx); err != nil {
                slog.Error("HTTP server shutdown error", "error", err)
            }
        }()
        
        // Shutdown worker pool
        wg.Add(1)
        go func() {
            defer wg.Done()
            if err := app.workerPool.GracefulShutdown(30 * time.Second); err != nil {
                slog.Error("Worker pool shutdown error", "error", err)
            }
        }()
        
        // Shutdown background services
        for _, service := range app.bgServices {
            wg.Add(1)
            go func(svc *BackgroundService) {
                defer wg.Done()
                if err := svc.GracefulShutdown(15 * time.Second); err != nil {
                    slog.Error("Background service shutdown error", "error", err)
                }
            }(service)
        }
        
        // Wait for all components to shutdown
        done := make(chan struct{})
        go func() {
            wg.Wait()
            close(done)
        }()
        
        select {
        case <-done:
            slog.Info("All components shut down gracefully")
        case <-ctx.Done():
            slog.Warn("Shutdown timeout reached, forcing exit")
        }
        
        // Close database connections
        if app.db != nil {
            app.db.Close()
            slog.Info("Database connections closed")
        }
        
        duration := time.Since(start)
        app.monitor.RecordShutdown(context.Background(), duration, ctx.Err() == nil)
        app.monitor.SetShutdownInProgress(context.Background(), false)
        
        slog.Info("Graceful shutdown completed", "duration", duration)
        close(app.shutdownChan)
    })
}

func (app *Application) shutdownHTTPServer(ctx context.Context) error {
    slog.Info("Shutting down HTTP server...")
    return app.server.Shutdown(ctx)
}

func (app *Application) healthCheck(w http.ResponseWriter, r *http.Request) {
    // Check if shutting down
    select {
    case <-app.shutdownChan:
        w.WriteHeader(http.StatusServiceUnavailable)
        w.Write([]byte("Shutting down"))
        return
    default:
    }
    
    // Basic health check
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}

func (app *Application) readinessCheck(w http.ResponseWriter, r *http.Request) {
    // Check if shutting down
    select {
    case <-app.shutdownChan:
        w.WriteHeader(http.StatusServiceUnavailable)
        w.Write([]byte("Not ready - shutting down"))
        return
    default:
    }
    
    // Check database connectivity
    if err := app.db.Ping(r.Context()); err != nil {
        w.WriteHeader(http.StatusServiceUnavailable)
        w.Write([]byte("Database unavailable"))
        return
    }
    
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Ready"))
}

func (app *Application) cleanupTask(ctx context.Context) error {
    // Perform cleanup operations
    slog.Debug("Running cleanup task")
    return nil
}

func (app *Application) metricsTask(ctx context.Context) error {
    // Collect and report metrics
    slog.Debug("Running metrics task")
    return nil
}

func main() {
    // Initialize structured logging
    slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
        Level: slog.LevelInfo,
    })))
    
    // Create and initialize application
    app := NewApplication()
    if err := app.Initialize(); err != nil {
        slog.Error("Application initialization failed", "error", err)
        os.Exit(1)
    }
    
    // Start application
    if err := app.Start(); err != nil {
        slog.Error("Application start failed", "error", err)
        os.Exit(1)
    }
    
    slog.Info("Application started successfully")
    
    // Wait for shutdown
    app.WaitForShutdown()
    
    slog.Info("Application exited")
}
```

## Best Practices Summary

### DO's

1. **Use Context for Cancellation**: Always use `context.Context` to signal cancellation
2. **Set Reasonable Timeouts**: Implement timeout contexts for all shutdown operations
3. **Handle Multiple Signals**: Listen for both SIGTERM and SIGINT
4. **Monitor Shutdown Progress**: Use metrics to track shutdown performance
5. **Test Shutdown Scenarios**: Include shutdown testing in your test suite
6. **Use Structured Logging**: Log shutdown progress with contextual information
7. **Close Resources in Reverse Order**: Close resources in the reverse order of initialization
8. **Implement Health Checks**: Provide health endpoints that reflect shutdown state

### DON'Ts

1. **Don't Block Forever**: Always use timeouts for shutdown operations
2. **Don't Ignore Cleanup**: Properly close all resources and connections
3. **Don't Handle Signals in Multiple Places**: Use a single signal handler
4. **Don't Forget Database Connections**: Ensure database connections are properly closed
5. **Don't Skip Error Handling**: Log and handle shutdown errors appropriately
6. **Don't Use os.Exit() Directly**: Allow graceful shutdown to complete first

### Checklist for Production Applications

- [ ] Signal handling implemented for SIGTERM and SIGINT
- [ ] HTTP server graceful shutdown with timeout
- [ ] Database connection draining implemented
- [ ] Background goroutines properly terminated
- [ ] Resource cleanup in correct order
- [ ] Health check endpoints reflect shutdown state
- [ ] Shutdown monitoring and metrics
- [ ] Container orchestration integration
- [ ] Comprehensive error handling
- [ ] Shutdown testing implemented

## Conclusion

Implementing graceful shutdown in Go applications requires careful consideration of signal handling, resource cleanup, and coordination between different application components. The patterns and examples provided in this document offer a comprehensive foundation for building robust, production-ready Go applications that can handle shutdown scenarios gracefully while maintaining data integrity and providing observability into the shutdown process.

The key to successful graceful shutdown is planning for it from the beginning of your application design, using appropriate libraries and patterns, and thoroughly testing shutdown scenarios in your CI/CD pipeline.

---

**Research Sources**:
- Go standard library documentation (os/signal, context, net/http)
- Industry best practices from major Go frameworks (Chi, Gin, Echo, Fiber)
- Container orchestration documentation (Kubernetes, Docker)
- Production patterns from large-scale Go applications
- Community libraries and patterns (oklog/run, vardius/shutdown, uber-go/fx)

**Last Updated**: 2025-08-27  
**Version**: 1.0  
**Author**: System Architect  
**Review Status**: Pending Technical Review