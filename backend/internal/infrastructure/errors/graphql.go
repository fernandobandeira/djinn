package errors

import (
	"context"
	"errors"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// GraphQLErrorPresenter creates a clean, strategy-based GraphQL error presenter
func GraphQLErrorPresenter(logger *slog.Logger) graphql.ErrorPresenterFunc {
	mapper := NewErrorMapper()
	
	return func(ctx context.Context, err error) *gqlerror.Error {
		// Create error context
		errorCtx := CreateErrorContext(ctx, logger)
		
		// Convert to gqlerror if needed
		var gqlErr *gqlerror.Error
		if !errors.As(err, &gqlErr) {
			gqlErr = &gqlerror.Error{
				Message: err.Error(),
			}
		}
		
		// Map the error using our clean strategy pattern
		mapping := mapper.MapError(err, errorCtx)
		
		// Apply the mapping
		gqlErr.Message = mapping.Message
		gqlErr.Extensions = mapping.Extensions
		
		// Log with appropriate level
		logger.Log(ctx, mapping.LogLevel, "GraphQL error",
			"correlation_id", errorCtx.CorrelationID,
			"error_code", string(mapping.Code),
			"error", err.Error(),
		)
		
		return gqlErr
	}
}