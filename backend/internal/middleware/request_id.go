package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const (
	// CorrelationIDKey is the context key for correlation ID
	CorrelationIDKey contextKey = "correlation_id"
)

// RequestID middleware adds a unique correlation ID to each request
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := r.Header.Get("X-Correlation-ID")
		if correlationID == "" {
			correlationID = uuid.New().String()
		}
		
		ctx := context.WithValue(r.Context(), CorrelationIDKey, correlationID)
		
		w.Header().Set("X-Correlation-ID", correlationID)
		
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetCorrelationID retrieves the correlation ID from context
func GetCorrelationID(ctx context.Context) string {
	if correlationID, ok := ctx.Value(CorrelationIDKey).(string); ok {
		return correlationID
	}
	// Fallback to request ID for backward compatibility
	if requestID, ok := ctx.Value(RequestIDKey).(string); ok {
		return requestID
	}
	return ""
}

// GetRequestID retrieves the request ID from context (deprecated, use GetCorrelationID)
// Maintained for backward compatibility
func GetRequestID(ctx context.Context) string {
	return GetCorrelationID(ctx)
}