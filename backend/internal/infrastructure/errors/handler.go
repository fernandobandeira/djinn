package errors

import (
	"context"
	"log/slog"

	ctxutil "github.com/fernandobandeira/djinn/backend/internal/infrastructure/context"
)

// GetErrorCode efficiently maps an error to its corresponding error code
func GetErrorCode(err error) ErrorCode {
	// Create minimal context for the mapper
	ctx := context.Background()
	logger := slog.Default()
	
	// Use the efficient error mapper
	mapping := MapError(err, ctx, logger)
	return mapping.Code
}

// getCorrelationID extracts correlation ID from context
func getCorrelationID(ctx context.Context) string {
	if id := ctxutil.GetCorrelationID(ctx); id != "" {
		return id
	}
	return "unknown"
}