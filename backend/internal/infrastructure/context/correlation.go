package context

import (
	"context"
)

type contextKey string

const (
	// CorrelationIDKey is the context key for correlation ID
	CorrelationIDKey contextKey = "correlation_id"
	// RequestIDKey is the context key for request ID (deprecated, for backward compatibility)
	RequestIDKey contextKey = "request_id"
)

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

// WithCorrelationID adds a correlation ID to the context
func WithCorrelationID(ctx context.Context, correlationID string) context.Context {
	return context.WithValue(ctx, CorrelationIDKey, correlationID)
}

// GetRequestID retrieves the request ID from context (deprecated, use GetCorrelationID)
// Maintained for backward compatibility
func GetRequestID(ctx context.Context) string {
	return GetCorrelationID(ctx)
}