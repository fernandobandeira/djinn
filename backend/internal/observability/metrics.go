package observability

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// HTTP metrics as defined in ADR
	HTTPDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "djinn_http_duration_seconds",
		Help:    "Duration of HTTP requests.",
		Buckets: prometheus.DefBuckets,
	}, []string{"method", "route", "status"})

	HTTPRequests = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "djinn_http_requests_total",
		Help: "Total number of HTTP requests.",
	}, []string{"method", "route", "status"})

	// Database metrics
	DBConnections = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "djinn_db_connections_active",
		Help: "Number of active database connections.",
	})

	DBQueryDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "djinn_db_query_duration_seconds",
		Help:    "Duration of database queries.",
		Buckets: []float64{0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 2, 5},
	}, []string{"operation", "table"})

	// Business metrics for receipt processing
	ReceiptProcessing = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "djinn_receipt_processing_seconds",
		Help:    "Time taken to process receipts.",
		Buckets: []float64{0.1, 0.5, 1, 2, 5, 10},
	})

	ReceiptProcessingTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "djinn_receipt_processing_total",
		Help: "Total number of receipts processed.",
	}, []string{"status", "category"})

	// Crash and error metrics
	PanicCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "djinn_panics_total",
		Help: "Total number of panic recoveries.",
	}, []string{"service", "handler"})

	CrashReports = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "djinn_crash_reports_total",
		Help: "Total number of crash reports.",
	}, []string{"severity", "category"})

	// Authentication metrics
	AuthAttempts = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "djinn_auth_attempts_total",
		Help: "Total number of authentication attempts.",
	}, []string{"method", "status"})

	// GraphQL specific metrics
	GraphQLDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "djinn_graphql_duration_seconds",
		Help:    "Duration of GraphQL operations.",
		Buckets: prometheus.DefBuckets,
	}, []string{"operation_type", "operation_name"})

	GraphQLErrors = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "djinn_graphql_errors_total",
		Help: "Total number of GraphQL errors.",
	}, []string{"operation_type", "error_type"})

	// Background job metrics
	BackgroundJobsProcessed = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "djinn_background_jobs_processed_total",
		Help: "Total number of background jobs processed.",
	}, []string{"job_type", "status"})

	BackgroundJobDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "djinn_background_job_duration_seconds",
		Help: "Duration of background job processing.",
	}, []string{"job_type"})

	// Memory and resource metrics
	AppMemoryUsage = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "djinn_app_memory_bytes",
		Help: "Application memory usage in bytes.",
	})

	GoroutinesCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "djinn_goroutines_count",
		Help: "Number of active goroutines.",
	})
)

// MetricsConfig holds configuration for Prometheus metrics
type MetricsConfig struct {
	Enabled bool
	Port    int
	Path    string
}

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	written    bool
}

func (rw *responseWriter) WriteHeader(code int) {
	if !rw.written {
		rw.statusCode = code
		rw.written = true
		rw.ResponseWriter.WriteHeader(code)
	}
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	if !rw.written {
		rw.WriteHeader(http.StatusOK)
	}
	return rw.ResponseWriter.Write(data)
}

// InitMetrics initializes Prometheus metrics endpoint
func InitMetrics(config MetricsConfig, logger *slog.Logger) (<-chan error, error) {
	metricsError := make(chan error, 1)
	
	if !config.Enabled {
		close(metricsError)
		return metricsError, nil
	}

	// Set default values
	if config.Port == 0 {
		config.Port = 9091
	}
	if config.Path == "" {
		config.Path = "/metrics"
	}

	// Create metrics server
	mux := http.NewServeMux()
	mux.Handle(config.Path, promhttp.Handler())

	// Create server with proper configuration
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start metrics server in background with proper error handling
	go func() {
		logger.Info("Starting metrics server",
			slog.Int("port", config.Port),
			slog.String("path", config.Path),
		)
		
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Metrics server failed",
				slog.String("error", err.Error()),
				slog.Int("port", config.Port),
			)
			// Send error but don't block
			select {
			case metricsError <- err:
			default:
			}
		}
	}()

	// Verify metrics server is accessible
	time.Sleep(100 * time.Millisecond)
	testURL := fmt.Sprintf("http://localhost:%d%s", config.Port, config.Path)
	client := &http.Client{Timeout: 1 * time.Second}
	if _, err := client.Get(testURL); err != nil {
		logger.Warn("Metrics server verification failed, continuing anyway",
			slog.String("url", testURL),
			slog.String("error", err.Error()),
		)
	} else {
		logger.Info("Metrics server verified and running",
			slog.String("url", testURL),
		)
	}

	return metricsError, nil
}

// HTTPMetricsMiddleware tracks HTTP metrics
func HTTPMetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Wrap response writer to capture status
		wrapped := &responseWriter{
			ResponseWriter: w,
			statusCode:     200,
		}

		// Process request
		next.ServeHTTP(wrapped, r)

		// Record metrics
		duration := time.Since(start).Seconds()
		statusStr := fmt.Sprintf("%d", wrapped.statusCode)

		HTTPDuration.WithLabelValues(r.Method, r.URL.Path, statusStr).Observe(duration)
		HTTPRequests.WithLabelValues(r.Method, r.URL.Path, statusStr).Inc()
	})
}

// RecordReceiptProcessing records receipt processing metrics
func RecordReceiptProcessing(ctx context.Context, duration time.Duration, success bool, category string) {
	ReceiptProcessing.Observe(duration.Seconds())

	status := "success"
	if !success {
		status = "failure"
	}
	ReceiptProcessingTotal.WithLabelValues(status, category).Inc()
}

// RecordDatabaseQuery records database query metrics
func RecordDatabaseQuery(operation, table string, duration time.Duration) {
	DBQueryDuration.WithLabelValues(operation, table).Observe(duration.Seconds())
}

// RecordAuthAttempt records authentication attempt metrics
func RecordAuthAttempt(method string, success bool) {
	status := "success"
	if !success {
		status = "failure"
	}
	AuthAttempts.WithLabelValues(method, status).Inc()
}

// RecordGraphQLOperation records GraphQL operation metrics
func RecordGraphQLOperation(operationType, operationName string, duration time.Duration, hasErrors bool) {
	GraphQLDuration.WithLabelValues(operationType, operationName).Observe(duration.Seconds())
	
	if hasErrors {
		GraphQLErrors.WithLabelValues(operationType, "execution").Inc()
	}
}

// RecordBackgroundJob records background job processing metrics
func RecordBackgroundJob(jobType string, duration time.Duration, success bool) {
	status := "success"
	if !success {
		status = "failure"
	}
	
	BackgroundJobsProcessed.WithLabelValues(jobType, status).Inc()
	BackgroundJobDuration.WithLabelValues(jobType).Observe(duration.Seconds())
}

// UpdateResourceMetrics updates application resource metrics
func UpdateResourceMetrics(memoryBytes uint64, goroutines int) {
	AppMemoryUsage.Set(float64(memoryBytes))
	GoroutinesCount.Set(float64(goroutines))
}