package resolver

// This file contains health check related GraphQL resolvers

import (
	"context"
	"fmt"
)

// Health is the resolver for the health field.
func (r *queryResolver) Health(ctx context.Context) (string, error) {
	// Check database connection
	if err := r.DB.Ping(ctx); err != nil {
		return "unhealthy", fmt.Errorf("database connection failed: %w", err)
	}
	return "healthy", nil
}