package observability

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"regexp"
	"runtime"
	"runtime/debug"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/posthog/posthog-go"

	ctxutil "github.com/fernandobandeira/djinn/backend/internal/infrastructure/context"
)

// Pre-compiled regex patterns for performance
var (
	emailRegex    = regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	jwtRegex      = regexp.MustCompile(`eyJ[a-zA-Z0-9_-]+\.[a-zA-Z0-9_-]+\.[a-zA-Z0-9_-]+`)
	apiKeyRegex   = regexp.MustCompile(`(api[_-]?key|apikey|api_secret|access_token|auth_token)["\s]*[:=]["\s]*[a-zA-Z0-9_\-]{20,}`)
	passwordRegex = regexp.MustCompile(`(password|passwd|pwd)["\s]*[:=]["\s]*[^"\s]+`)
	ccRegex       = regexp.MustCompile(`\b(?:\d[ -]*?){13,16}\b`)
	ssnRegex      = regexp.MustCompile(`\b\d{3}-\d{2}-\d{4}\b`)
	dbRegex       = regexp.MustCompile(`(postgres|mysql|mongodb|redis):\/\/[^@]+@[^\s]+`)
)

// CrashReport represents a crash report with context as defined in ADR
type CrashReport struct {
	ID            string                 `json:"id"`
	Timestamp     time.Time              `json:"timestamp"`
	CorrelationID string                 `json:"correlation_id"`
	UserID        string                 `json:"user_id,omitempty"`
	SessionID     string                 `json:"session_id,omitempty"`
	Severity      string                 `json:"severity"` // critical, error, warning
	Category      string                 `json:"category"` // panic, error, timeout
	Message       string                 `json:"message"`
	StackTrace    string                 `json:"stack_trace,omitempty"`
	Context       map[string]interface{} `json:"context"`
	Breadcrumbs   []Breadcrumb           `json:"breadcrumbs,omitempty"`
	SystemInfo    SystemInfo             `json:"system_info"`
}

// SystemInfo contains system information for crash reports
type SystemInfo struct {
	GoVersion    string `json:"go_version"`
	NumGoroutine int    `json:"num_goroutine"`
	MemStats     struct {
		Alloc      uint64 `json:"alloc"`
		TotalAlloc uint64 `json:"total_alloc"`
		Sys        uint64 `json:"sys"`
		NumGC      uint32 `json:"num_gc"`
	} `json:"mem_stats"`
}

// CrashReporter handles crash reporting
type CrashReporter struct {
	analytics *AnalyticsTracker
	logger    *slog.Logger
	enabled   bool
}

// NewCrashReporter creates a new crash reporter
func NewCrashReporter(analytics *AnalyticsTracker, logger *slog.Logger, enabled bool) *CrashReporter {
	return &CrashReporter{
		analytics: analytics,
		logger:    logger,
		enabled:   enabled,
	}
}

// PanicRecoveryMiddleware recovers from panics and reports crashes
// Integrates with error-handling-ADR patterns
func (cr *CrashReporter) PanicRecoveryMiddleware(serviceName string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					// Capture panic details
					stackTrace := string(debug.Stack())
					correlationID := ctxutil.GetCorrelationID(r.Context())
					if correlationID == "" {
						correlationID = uuid.New().String()
					}

					// Log with slog (from error-handling-ADR)
					cr.logger.Error("Panic recovered",
						slog.String("correlation_id", correlationID),
						slog.String("service", serviceName),
						slog.String("path", r.URL.Path),
						slog.Any("panic", err),
						slog.String("stack_trace", stackTrace),
					)

					// Track panic metric
					PanicCounter.WithLabelValues(serviceName, r.URL.Path).Inc()

					// Create crash report with sanitized data
					report := CrashReport{
						ID:            uuid.New().String(),
						Timestamp:     time.Now().UTC(),
						CorrelationID: correlationID,
						UserID:        GetUserID(r.Context()),
						SessionID:     GetSessionID(r.Context()),
						Severity:      "critical",
						Category:      "panic",
						Message:       sanitizeMessage(fmt.Sprintf("%v", err)),
						StackTrace:    stackTrace,
						Context: sanitizeContext(map[string]interface{}{
							"method":     r.Method,
							"path":       r.URL.Path,
							"user_agent": r.UserAgent(),
							"remote_ip":  r.RemoteAddr,
							"service":    serviceName,
						}),
						SystemInfo: collectSystemInfo(),
					}

					// Add breadcrumbs if available
					if cr.analytics != nil {
						report.Breadcrumbs = cr.analytics.GetBreadcrumbs(correlationID)
					}

					// Send crash report
					go cr.SendCrashReport(report)

					// Return error response
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(`{"error":"Internal server error","correlation_id":"` + correlationID + `"}`))
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

// SendCrashReport sends crash report to monitoring backend
func (cr *CrashReporter) SendCrashReport(report CrashReport) {
	if !cr.enabled {
		return
	}

	// Track in metrics
	CrashReports.WithLabelValues(report.Severity, report.Category).Inc()

	// Send to PostHog if analytics is enabled
	if cr.analytics != nil && cr.analytics.enabled {
		properties := posthog.NewProperties().
			Set("$exception_message", report.Message).
			Set("$exception_type", report.Category).
			Set("$exception_stack_trace", report.StackTrace).
			Set("correlation_id", report.CorrelationID).
			Set("severity", report.Severity).
			Set("system_info", report.SystemInfo)

		// Add breadcrumbs if present
		if len(report.Breadcrumbs) > 0 {
			breadcrumbData := make([]map[string]interface{}, len(report.Breadcrumbs))
			for i, crumb := range report.Breadcrumbs {
				breadcrumbData[i] = map[string]interface{}{
					"timestamp": crumb.Timestamp.Format(time.RFC3339),
					"category":  crumb.Category,
					"message":   crumb.Message,
					"data":      crumb.Data,
				}
			}
			properties.Set("breadcrumbs", breadcrumbData)
		}

		distinctID := report.UserID
		if distinctID == "" {
			distinctID = report.CorrelationID
		}

		_ = cr.analytics.client.Enqueue(posthog.Capture{
			DistinctId: distinctID,
			Event:      "$exception",
			Properties: properties,
		})
	}

	// Log for local debugging with sanitized data
	cr.logger.Error("Crash report",
		slog.String("report_id", report.ID),
		slog.String("correlation_id", report.CorrelationID),
		slog.String("severity", report.Severity),
		slog.String("category", report.Category),
		slog.String("message", report.Message), // Message is already sanitized
	)
}

// ReportError reports non-panic errors for tracking
func (cr *CrashReporter) ReportError(ctx context.Context, err error, category string, severity string) {
	if !cr.enabled || err == nil {
		return
	}

	correlationID := ctxutil.GetCorrelationID(ctx)
	userID := GetUserID(ctx)

	report := CrashReport{
		ID:            uuid.New().String(),
		Timestamp:     time.Now().UTC(),
		CorrelationID: correlationID,
		UserID:        userID,
		SessionID:     GetSessionID(ctx),
		Severity:      severity,
		Category:      category,
		Message:       err.Error(),
		Context: map[string]interface{}{
			"error_type": fmt.Sprintf("%T", err),
		},
		SystemInfo: collectSystemInfo(),
	}

	// Add breadcrumbs
	if cr.analytics != nil {
		report.Breadcrumbs = cr.analytics.GetBreadcrumbs(correlationID)
	}

	go cr.SendCrashReport(report)
}

// ReportTimeout reports timeout errors
func (cr *CrashReporter) ReportTimeout(ctx context.Context, operation string, duration time.Duration) {
	if !cr.enabled {
		return
	}

	report := CrashReport{
		ID:            uuid.New().String(),
		Timestamp:     time.Now().UTC(),
		CorrelationID: ctxutil.GetCorrelationID(ctx),
		UserID:        GetUserID(ctx),
		SessionID:     GetSessionID(ctx),
		Severity:      "error",
		Category:      "timeout",
		Message:       fmt.Sprintf("Operation '%s' timed out after %v", operation, duration),
		Context: map[string]interface{}{
			"operation": operation,
			"duration":  duration.Seconds(),
		},
		SystemInfo: collectSystemInfo(),
	}

	if cr.analytics != nil {
		report.Breadcrumbs = cr.analytics.GetBreadcrumbs(report.CorrelationID)
	}

	go cr.SendCrashReport(report)
}

// collectSystemInfo gathers system information for crash reports
func collectSystemInfo() SystemInfo {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return SystemInfo{
		GoVersion:    runtime.Version(),
		NumGoroutine: runtime.NumGoroutine(),
		MemStats: struct {
			Alloc      uint64 `json:"alloc"`
			TotalAlloc uint64 `json:"total_alloc"`
			Sys        uint64 `json:"sys"`
			NumGC      uint32 `json:"num_gc"`
		}{
			Alloc:      m.Alloc,
			TotalAlloc: m.TotalAlloc,
			Sys:        m.Sys,
			NumGC:      m.NumGC,
		},
	}
}

// Context key types
type contextKey string

const (
	userIDKey    contextKey = "user_id"
	sessionIDKey contextKey = "session_id"
)


// GetUserID gets user ID from context
func GetUserID(ctx context.Context) string {
	if id, ok := ctx.Value(userIDKey).(string); ok {
		return id
	}
	return ""
}

// GetSessionID gets session ID from context
func GetSessionID(ctx context.Context) string {
	if id, ok := ctx.Value(sessionIDKey).(string); ok {
		return id
	}
	return ""
}

// WithCorrelationID adds correlation ID to context
func WithCorrelationID(ctx context.Context, id string) context.Context {
	return ctxutil.WithCorrelationID(ctx, id)
}

// WithUserID adds user ID to context
func WithUserID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, userIDKey, id)
}

// WithSessionID adds session ID to context
func WithSessionID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, sessionIDKey, id)
}

// sanitizeMessage removes potentially sensitive data from messages
func sanitizeMessage(msg string) string {
	// Use pre-compiled regex patterns for better performance
	msg = emailRegex.ReplaceAllString(msg, "[EMAIL]")
	msg = jwtRegex.ReplaceAllString(msg, "[JWT_TOKEN]")
	msg = apiKeyRegex.ReplaceAllString(msg, "$1=[REDACTED]")
	msg = passwordRegex.ReplaceAllString(msg, "$1=[REDACTED]")
	msg = ccRegex.ReplaceAllString(msg, "[CREDIT_CARD]")
	msg = ssnRegex.ReplaceAllString(msg, "[SSN]")
	msg = dbRegex.ReplaceAllString(msg, "$1://[REDACTED]")
	
	return msg
}

// sanitizeContext removes sensitive data from context maps
func sanitizeContext(ctx map[string]interface{}) map[string]interface{} {
	sanitized := make(map[string]interface{})
	sensitiveKeys := []string{
		"password", "passwd", "pwd",
		"token", "secret", "key",
		"authorization", "auth",
		"credit_card", "cc", "cvv",
		"ssn", "social_security",
	}
	
	for k, v := range ctx {
		// Check if key contains sensitive words
		keyLower := strings.ToLower(k)
		isSensitive := false
		for _, sensitive := range sensitiveKeys {
			if strings.Contains(keyLower, sensitive) {
				isSensitive = true
				break
			}
		}
		
		if isSensitive {
			sanitized[k] = "[REDACTED]"
		} else {
			// Recursively sanitize nested maps
			if nestedMap, ok := v.(map[string]interface{}); ok {
				sanitized[k] = sanitizeContext(nestedMap)
			} else if str, ok := v.(string); ok {
				sanitized[k] = sanitizeMessage(str)
			} else {
				sanitized[k] = v
			}
		}
	}
	
	return sanitized
}