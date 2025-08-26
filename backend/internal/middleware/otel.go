package middleware

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// OTelHTTPMiddleware creates OpenTelemetry HTTP middleware using the contrib package
func OTelHTTPMiddleware(serviceName string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		// otelhttp handles all the complexity of tracing HTTP requests
		// including trace propagation, span creation, and attribute recording
		return otelhttp.NewHandler(next, serviceName,
			otelhttp.WithMessageEvents(otelhttp.ReadEvents, otelhttp.WriteEvents),
		)
	}
}