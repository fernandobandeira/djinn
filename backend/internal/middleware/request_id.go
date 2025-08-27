package middleware

import (
	"context"
	"net/http"

	ctxutil "github.com/fernandobandeira/djinn/backend/internal/infrastructure/context"
	"github.com/google/uuid"
)

// RequestID middleware adds a unique correlation ID to each request
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := r.Header.Get("X-Correlation-ID")
		if correlationID == "" {
			correlationID = uuid.New().String()
		}
		
		ctx := ctxutil.WithCorrelationID(r.Context(), correlationID)
		
		w.Header().Set("X-Correlation-ID", correlationID)
		
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetCorrelationID retrieves the correlation ID from context
func GetCorrelationID(ctx context.Context) string {
	return ctxutil.GetCorrelationID(ctx)
}

// GetRequestID retrieves the request ID from context (deprecated, use GetCorrelationID)
func GetRequestID(ctx context.Context) string {
	return ctxutil.GetCorrelationID(ctx)
}