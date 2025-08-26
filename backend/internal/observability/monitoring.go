package observability

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"regexp"
	"runtime"
	"time"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

// MonitoringService integrates all observability components as per ADR
type MonitoringService struct {
	Analytics     *AnalyticsTracker
	CrashReporter *CrashReporter
	config        *MonitoringConfig
	logger        *slog.Logger
}

// MonitoringConfig holds all monitoring configuration
type MonitoringConfig struct {
	// Metrics configuration
	MetricsEnabled bool
	MetricsPort    int

	// Tracing configuration
	TracingEnabled    bool
	TracingSampleRate float64
	OTLPEndpoint      string

	// Analytics configuration
	AnalyticsEnabled bool
	PostHogAPIKey    string

	// Service information
	ServiceName    string
	ServiceVersion string
	Environment    string
}

// NewMonitoringService creates a new integrated monitoring service
func NewMonitoringService(ctx context.Context, config *MonitoringConfig, logger *slog.Logger) (*MonitoringService, func(context.Context) error, error) {
	var shutdownFuncs []func(context.Context) error

	// Initialize tracing with fallback
	tracingConfig := TracingConfig{
		ServiceName:    config.ServiceName,
		ServiceVersion: config.ServiceVersion,
		Environment:    config.Environment,
		OTLPEndpoint:   config.OTLPEndpoint,
		Enabled:        config.TracingEnabled,
		SampleRate:     config.TracingSampleRate,
	}

	tracingShutdown, err := InitTracer(ctx, tracingConfig)
	if err != nil {
		// Log warning but continue without tracing
		logger.Warn("Tracing initialization failed, continuing without distributed tracing",
			slog.String("error", err.Error()),
			slog.String("endpoint", config.OTLPEndpoint),
		)
		// Provide no-op shutdown function
		tracingShutdown = func(context.Context) error { return nil }
	} else if config.TracingEnabled {
		// Verify tracing is working
		if err := verifyTracing(ctx); err != nil {
			logger.Warn("Tracing verification failed",
				slog.String("error", err.Error()),
			)
		} else {
			logger.Info("Tracing initialized and verified",
				slog.String("endpoint", config.OTLPEndpoint),
				slog.Float64("sample_rate", config.TracingSampleRate),
			)
		}
	}
	if tracingShutdown != nil {
		shutdownFuncs = append(shutdownFuncs, tracingShutdown)
	}

	// Initialize metrics
	metricsConfig := MetricsConfig{
		Enabled: config.MetricsEnabled,
		Port:    config.MetricsPort,
		Path:    "/metrics",
	}

	metricsErrorChan, err := InitMetrics(metricsConfig, logger)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to initialize metrics: %w", err)
	}
	
	// Monitor for metrics server errors
	go func() {
		if err := <-metricsErrorChan; err != nil {
			logger.Error("Metrics server error", slog.String("error", err.Error()))
		}
	}()

	// Start resource metrics collector
	if config.MetricsEnabled {
		go collectResourceMetrics()
	}

	// Initialize analytics
	analyticsConfig := AnalyticsConfig{
		APIKey:      config.PostHogAPIKey,
		Enabled:     config.AnalyticsEnabled,
		Environment: config.Environment,
	}

	analytics, err := NewAnalyticsTracker(analyticsConfig, logger)
	if err != nil {
		logger.Warn("Failed to initialize analytics, continuing without analytics",
			slog.String("error", err.Error()),
		)
		analytics = &AnalyticsTracker{enabled: false}
	}
	if analytics != nil && analytics.client != nil {
		shutdownFuncs = append(shutdownFuncs, func(ctx context.Context) error {
			return analytics.Close()
		})
	}

	// Initialize crash reporter
	crashReporter := NewCrashReporter(analytics, logger, config.AnalyticsEnabled)

	// Combined shutdown function
	shutdown := func(ctx context.Context) error {
		var firstErr error
		for _, fn := range shutdownFuncs {
			if err := fn(ctx); err != nil && firstErr == nil {
				firstErr = err
			}
		}
		return firstErr
	}

	return &MonitoringService{
		Analytics:     analytics,
		CrashReporter: crashReporter,
		config:        config,
		logger:        logger,
	}, shutdown, nil
}

// HTTPMiddleware creates a comprehensive HTTP middleware chain
func (ms *MonitoringService) HTTPMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Generate and validate correlation ID
			correlationID := r.Header.Get("X-Correlation-ID")
			if correlationID == "" {
				correlationID = uuid.New().String()
			} else {
				correlationID = validateCorrelationID(correlationID)
			}

			// Add correlation ID to context and response
			ctx := WithCorrelationID(r.Context(), correlationID)
			w.Header().Set("X-Correlation-ID", correlationID)

			// Extract user ID if present (from auth middleware) - capture before processing
			userID := ""
			if uid := r.Header.Get("X-User-ID"); uid != "" {
				userID = uid
				ctx = WithUserID(ctx, uid)
			}

			// Extract session ID if present - capture before processing
			if sid := r.Header.Get("X-Session-ID"); sid != "" {
				ctx = WithSessionID(ctx, sid)
			}

			// Update request with new context
			r = r.WithContext(ctx)

			// Add breadcrumb for request
			if ms.Analytics != nil {
				ms.Analytics.AddBreadcrumb(correlationID, Breadcrumb{
					Timestamp: time.Now(),
					Category:  "http",
					Message:   fmt.Sprintf("%s %s", r.Method, r.URL.Path),
					Data: map[string]interface{}{
						"user_agent": r.UserAgent(),
						"referer":    r.Referer(),
						"remote_ip":  r.RemoteAddr,
					},
				})
			}

			// Start timing
			start := time.Now()

			// Wrap response writer to capture status
			wrapped := &responseWriter{
				ResponseWriter: w,
				statusCode:     200,
			}

			// Apply panic recovery
			panicHandler := ms.CrashReporter.PanicRecoveryMiddleware(ms.config.ServiceName)
			panicHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Capture request details before processing (to avoid race conditions)
				method := r.Method
				path := r.URL.Path
				
				// Set up deferred metrics recording to avoid race conditions
				defer func() {
					// Record metrics after request completes (safe in defer)
					duration := time.Since(start).Seconds()
					statusStr := fmt.Sprintf("%d", wrapped.statusCode)

					HTTPDuration.WithLabelValues(method, path, statusStr).Observe(duration)
					HTTPRequests.WithLabelValues(method, path, statusStr).Inc()

					// Track user request if user ID was captured
					if userID != "" && ms.Analytics != nil {
						ms.Analytics.TrackRequest(userID, r, wrapped.statusCode, duration)
					}

					// Log slow requests
					if duration > 1.0 {
						ms.logger.Warn("Slow request detected",
							slog.String("correlation_id", correlationID),
							slog.String("method", method),
							slog.String("path", path),
							slog.Float64("duration", duration),
							slog.Int("status", wrapped.statusCode),
						)
					}

					// Add final breadcrumb
					if ms.Analytics != nil {
						ms.Analytics.AddBreadcrumb(correlationID, Breadcrumb{
							Timestamp: time.Now(),
							Category:  "http_response",
							Message:   fmt.Sprintf("Response %d", wrapped.statusCode),
							Data: map[string]interface{}{
								"duration_ms": duration * 1000,
							},
						})
					}
				}()
				
				// Process request
				next.ServeHTTP(wrapped, r)
			})).ServeHTTP(w, r)
		})
	}
}

// TrackReceiptProcessing tracks receipt processing metrics and analytics
func (ms *MonitoringService) TrackReceiptProcessing(ctx context.Context, duration time.Duration, success bool, details map[string]interface{}) {
	// Record metrics
	RecordReceiptProcessing(ctx, duration, success, details["category"].(string))

	// Track analytics
	if userID := GetUserID(ctx); userID != "" && ms.Analytics != nil {
		ms.Analytics.TrackReceiptProcessing(ctx, userID, duration, success, details)
	}

	// Report if timeout
	if duration > 10*time.Second && ms.CrashReporter != nil {
		ms.CrashReporter.ReportTimeout(ctx, "receipt_processing", duration)
	}
}

// TrackAuthAttempt tracks authentication attempts
func (ms *MonitoringService) TrackAuthAttempt(ctx context.Context, method string, success bool, properties map[string]interface{}) {
	// Record metrics
	RecordAuthAttempt(method, success)

	// Track analytics
	if userID := GetUserID(ctx); userID != "" && ms.Analytics != nil {
		event := "user_login"
		if method == "signup" {
			event = "user_signup_completed"
		}
		ms.Analytics.TrackAuthEvent(userID, event, properties)
	}
}

// TrackGraphQLOperation tracks GraphQL operations
func (ms *MonitoringService) TrackGraphQLOperation(ctx context.Context, operationType, operationName string, duration time.Duration, hasErrors bool) {
	RecordGraphQLOperation(operationType, operationName, duration, hasErrors)

	// Add breadcrumb
	if ms.Analytics != nil {
		correlationID := GetCorrelationID(ctx)
		ms.Analytics.AddBreadcrumb(correlationID, Breadcrumb{
			Timestamp: time.Now(),
			Category:  "graphql",
			Message:   fmt.Sprintf("%s %s", operationType, operationName),
			Data: map[string]interface{}{
				"duration_ms": duration.Milliseconds(),
				"has_errors":  hasErrors,
			},
		})
	}
}

// TrackBackgroundJob tracks background job processing
func (ms *MonitoringService) TrackBackgroundJob(ctx context.Context, jobType string, duration time.Duration, success bool) {
	RecordBackgroundJob(jobType, duration, success)

	// Report if failed
	if !success && ms.CrashReporter != nil {
		ms.CrashReporter.ReportError(ctx, fmt.Errorf("background job failed"), jobType, "error")
	}
}

// verifyTracing verifies that tracing is working
func verifyTracing(ctx context.Context) error {
	tracer := otel.Tracer("verification")
	ctx, span := tracer.Start(ctx, "tracing-verification-test")
	defer span.End()
	
	span.SetAttributes(
		attribute.Bool("test", true),
		attribute.String("purpose", "startup_verification"),
	)
	
	// If we got here without panic, tracing is working
	return nil
}

// validateCorrelationID validates and sanitizes correlation ID
func validateCorrelationID(id string) string {
	// UUID pattern
	uuidPattern := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`)
	
	if uuidPattern.MatchString(id) {
		return id
	}
	
	// If not valid UUID, generate new one
	return uuid.New().String()
}

// collectResourceMetrics periodically collects resource metrics
func collectResourceMetrics() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		
		UpdateResourceMetrics(m.Alloc, runtime.NumGoroutine())
	}
}