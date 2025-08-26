package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fernandobandeira/djinn/backend/internal/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// TracingMiddleware adds OpenTelemetry tracing to HTTP requests
func TracingMiddleware() func(http.Handler) http.Handler {
	tracer := observability.GetTracer("http")
	
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract trace context from incoming request
			ctx := propagation.TraceContext{}.Extract(r.Context(), propagation.HeaderCarrier(r.Header))
			
			// Start a new span for this request
			spanName := fmt.Sprintf("%s %s", r.Method, r.URL.Path)
			ctx, span := observability.StartSpan(ctx, tracer, spanName,
				trace.WithSpanKind(trace.SpanKindServer),
				trace.WithAttributes(
					attribute.String("http.method", r.Method),
					attribute.String("http.target", r.URL.String()),
					attribute.String("http.route", r.URL.Path),
					attribute.String("http.scheme", r.URL.Scheme),
					attribute.String("net.host.name", r.Host),
					attribute.String("http.user_agent", r.UserAgent()),
					attribute.Int64("http.request_content_length", r.ContentLength),
					attribute.String("net.peer.name", r.RemoteAddr),
				),
			)
			defer span.End()
			
			// Create wrapped response writer to capture status code
			wrapped := &responseWriter{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
				written:        0,
				startTime:      time.Now(),
			}
			
			// Pass request with trace context to next handler
			next.ServeHTTP(wrapped, r.WithContext(ctx))
			
			// Record response attributes
			span.SetAttributes(
				attribute.Int("http.status_code", wrapped.statusCode),
				attribute.Int64("http.response_content_length", wrapped.written),
				attribute.Float64("http.request.duration_ms", 
					float64(time.Since(wrapped.startTime).Milliseconds())),
			)
			
			// Set span status based on HTTP status code
			if wrapped.statusCode >= 400 {
				span.SetStatus(codes.Error, fmt.Sprintf("HTTP %d", wrapped.statusCode))
			} else {
				span.SetStatus(codes.Ok, "")
			}
			
			// Add request ID if present
			if reqID := GetRequestID(ctx); reqID != "" {
				span.SetAttributes(attribute.String("request.id", reqID))
			}
		})
	}
}

// responseWriter wraps http.ResponseWriter to capture status code and bytes written
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	written    int64
	startTime  time.Time
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(b)
	rw.written += int64(n)
	return n, err
}

// TracingGraphQLMiddleware adds tracing specifically for GraphQL operations
func TracingGraphQLMiddleware() func(http.Handler) http.Handler {
	tracer := observability.GetTracer("graphql")
	
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/api/graphql" && r.Method == http.MethodPost {
				ctx := r.Context()
				
				// Start a GraphQL-specific span
				ctx, span := observability.StartSpan(ctx, tracer, "GraphQL Operation",
					trace.WithSpanKind(trace.SpanKindServer),
				)
				defer span.End()
				
				// The actual operation details will be added by the GraphQL handler
				span.SetAttributes(
					attribute.String("graphql.endpoint", r.URL.Path),
					attribute.String("graphql.transport", "HTTP POST"),
				)
				
				r = r.WithContext(ctx)
			}
			
			next.ServeHTTP(w, r)
		})
	}
}