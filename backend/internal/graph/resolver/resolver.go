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
	DB     *database.DB
	Logger *slog.Logger
}

// NewResolver creates a new GraphQL resolver
func NewResolver(db *database.DB, logger *slog.Logger) *Resolver {
	return &Resolver{
		DB:     db,
		Logger: logger,
	}
}