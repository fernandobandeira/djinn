package resolver

import (
	"context"
	"log/slog"

	db "github.com/fernandobandeira/djinn/backend/internal/database/generated"
	"github.com/jackc/pgx/v5/pgtype"
)

// QueriesInterface defines the database operations we need
// This allows us to inject mocks for testing
type QueriesInterface interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	GetUser(ctx context.Context, id pgtype.UUID) (db.User, error)
	GetUserByFirebaseUID(ctx context.Context, firebaseUid string) (db.User, error)
	UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error)
	DeleteUser(ctx context.Context, id pgtype.UUID) error
	ListUsers(ctx context.Context, arg db.ListUsersParams) ([]db.User, error)
	GetUsersByIDs(ctx context.Context, ids []pgtype.UUID) ([]db.User, error)
}

// PingableDB defines the minimal database interface needed for health checks
type PingableDB interface {
	Ping(ctx context.Context) error
}

// LoggerInterface defines the logger operations we need
type LoggerInterface interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any) 
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

// Ensure our implementations satisfy the interfaces
var _ QueriesInterface = (*db.Queries)(nil)
var _ LoggerInterface = (*slog.Logger)(nil)