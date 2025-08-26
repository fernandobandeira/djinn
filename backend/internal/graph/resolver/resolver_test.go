package resolver

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolverMethods(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelError}))
	
	resolver := &Resolver{
		Logger: logger,
	}
	
	// Test Mutation method
	mutation := resolver.Mutation()
	assert.NotNil(t, mutation)
	assert.IsType(t, &mutationResolver{}, mutation)
	
	// Test Query method
	query := resolver.Query()
	assert.NotNil(t, query)
	assert.IsType(t, &queryResolver{}, query)
	
	// User resolvers are now part of Query/Mutation resolvers with follow-schema layout
}

func TestNewResolver(t *testing.T) {
	// This test is covered by Wire dependency injection integration tests
	// The NewResolver function now requires all command and query handlers
	// which are properly tested in their respective test files
	t.Skip("NewResolver is tested through Wire integration")
}

