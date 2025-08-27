package resolver

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	ctx := context.Background()
	
	// Create logger
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelError}))
	
	// Create resolver
	resolver := &queryResolver{
		Resolver: &Resolver{
			Logger: logger,
		},
	}
	
	// Execute
	result, err := resolver.Health(ctx)
	
	// Assert - health check should always return "healthy" without checking database
	assert.Equal(t, "healthy", result)
	assert.NoError(t, err)
}