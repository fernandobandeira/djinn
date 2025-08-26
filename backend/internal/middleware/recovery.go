package middleware

import (
	"log/slog"
	"net/http"
	"runtime/debug"
)

// Recovery middleware recovers from panics and logs them
func Recovery(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					requestID := GetRequestID(r.Context())
					
					logger.Error("panic recovered",
						"error", err,
						"request_id", requestID,
						"stack", string(debug.Stack()),
					)
					
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(`{"error":"Internal Server Error"}`))
				}
			}()
			
			next.ServeHTTP(w, r)
		})
	}
}