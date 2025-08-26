package middleware

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"

	apperrors "github.com/fernandobandeira/djinn/backend/internal/infrastructure/errors"
)

// Recovery middleware recovers from panics and logs them
func Recovery(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					correlationID := GetCorrelationID(r.Context())
					
					// Log the panic with full stack trace
					logger.Error("panic recovered",
						"error", err,
						"correlation_id", correlationID,
						"path", r.URL.Path,
						"method", r.Method,
						"stack", string(debug.Stack()),
					)
					
					// Create standardized error response
					errResp := apperrors.NewErrorResponse(
						apperrors.CodeInternalError,
						"An unexpected error occurred",
						correlationID,
					)
					
					// Write response
					w.Header().Set("Content-Type", "application/json")
					w.Header().Set("X-Correlation-ID", correlationID)
					w.WriteHeader(http.StatusInternalServerError)
					
					if encErr := json.NewEncoder(w).Encode(errResp); encErr != nil {
						logger.Error("failed to encode panic error response",
							"correlation_id", correlationID,
							"encoding_error", encErr.Error(),
							"original_panic", fmt.Sprintf("%v", err),
						)
					}
				}
			}()
			
			next.ServeHTTP(w, r)
		})
	}
}