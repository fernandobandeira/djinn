package errors

import (
	"context"
	"errors"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// GraphQLErrorPresenter creates a clean, efficient GraphQL error presenter
func GraphQLErrorPresenter(logger *slog.Logger) graphql.ErrorPresenterFunc {
	return func(ctx context.Context, err error) *gqlerror.Error {
		// Convert to gqlerror if needed
		var gqlErr *gqlerror.Error
		if !errors.As(err, &gqlErr) {
			gqlErr = &gqlerror.Error{
				Message: err.Error(),
			}
		}
		
		// Map the error using efficient type switching
		mapping := MapError(err, ctx, logger)
		
		// Apply the mapping
		gqlErr.Message = mapping.Message
		gqlErr.Extensions = mapping.Extensions
		
		// Log with appropriate level and structured context
		logger.Log(ctx, mapping.LogLevel, "GraphQL error",
			"error_code", string(mapping.Code),
			"error_message", err.Error(),
		)
		
		return gqlErr
	}
}