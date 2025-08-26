package resolver

import (
	"github.com/fernandobandeira/djinn/backend/internal/database"
	"log/slog"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the GraphQL resolver root
type Resolver struct {
	DB       DatabaseInterface
	Queries  QueriesInterface
	Logger   LoggerInterface
}

// NewResolver creates a new GraphQL resolver
func NewResolver(db *database.DB, logger *slog.Logger) *Resolver {
	wrapper := NewDatabaseWrapper(db)
	return &Resolver{
		DB:      wrapper,
		Queries: db.Queries,
		Logger:  logger,
	}
}

// NewResolverWithInterfaces creates a resolver with injected interfaces for testing
func NewResolverWithInterfaces(queries QueriesInterface, logger LoggerInterface) *Resolver {
	return &Resolver{
		Queries: queries,
		Logger:  logger,
	}
}