package resolver

import (
	"context"

	"github.com/fernandobandeira/djinn/backend/internal/database"
)

// DatabaseWrapper wraps the database.DB to provide QueriesInterface
type DatabaseWrapper struct {
	*database.DB
}

// GetQueries returns the Queries as QueriesInterface
func (d *DatabaseWrapper) GetQueries() QueriesInterface {
	return d.Queries
}

// Ping checks database connectivity
func (d *DatabaseWrapper) Ping(ctx context.Context) error {
	return d.DB.Ping(ctx)
}

// NewDatabaseWrapper creates a new DatabaseWrapper
func NewDatabaseWrapper(db *database.DB) *DatabaseWrapper {
	return &DatabaseWrapper{DB: db}
}