package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

const (
	// SlowRequestThreshold defines when a request is considered slow
	SlowRequestThreshold = 5 * time.Second
)

// Logger middleware logs HTTP requests with structured logging
func Logger(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			
			defer func() {
				duration := time.Since(start)
				correlationID := GetCorrelationID(r.Context())
				
				// Determine log level based on response status and duration
				logLevel := slog.LevelInfo
				if ww.Status() >= 500 {
					logLevel = slog.LevelError
				} else if ww.Status() >= 400 {
					logLevel = slog.LevelWarn
				} else if duration > SlowRequestThreshold {
					logLevel = slog.LevelWarn
				}
				
				// Build log attributes
				attrs := []any{
					"method", r.Method,
					"path", r.URL.Path,
					"status", ww.Status(),
					"bytes", ww.BytesWritten(),
					"duration", duration,
					"duration_ms", duration.Milliseconds(),
					"correlation_id", correlationID,
					"remote_addr", r.RemoteAddr,
					"user_agent", r.UserAgent(),
				}
				
				// Add slow request indicator
				if duration > SlowRequestThreshold {
					attrs = append(attrs, "slow_request", true)
					attrs = append(attrs, "threshold_ms", SlowRequestThreshold.Milliseconds())
				}
				
				// Add query parameters if present
				if r.URL.RawQuery != "" {
					attrs = append(attrs, "query", r.URL.RawQuery)
				}
				
				// Log with appropriate level
				switch logLevel {
				case slog.LevelError:
					logger.Error("http request", attrs...)
				case slog.LevelWarn:
					logger.Warn("http request", attrs...)
				default:
					logger.Info("http request", attrs...)
				}
			}()
			
			next.ServeHTTP(ww, r)
		})
	}
}