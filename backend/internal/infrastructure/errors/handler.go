package errors

import (
	"context"
	"log/slog"

	ctxutil "github.com/fernandobandeira/djinn/backend/internal/infrastructure/context"
)

// GetErrorCode maps an error to its corresponding error code using the strategy pattern
func GetErrorCode(err error) ErrorCode {
	// Create a minimal context for the mapper
	ctx := context.Background()
	logger := slog.Default()
	errorCtx := CreateErrorContext(ctx, logger)
	
	// Use the error mapper to get the code
	mapper := NewErrorMapper()
	mapping := mapper.MapError(err, errorCtx)
	return mapping.Code
}

// getCorrelationID extracts correlation ID from context
func getCorrelationID(ctx context.Context) string {
	if id := ctxutil.GetCorrelationID(ctx); id != "" {
		return id
	}
	return "unknown"
}