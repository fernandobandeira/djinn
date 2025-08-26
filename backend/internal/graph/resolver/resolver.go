package resolver

import (
	commandUser "github.com/fernandobandeira/djinn/backend/internal/application/command/user"
	queryUser "github.com/fernandobandeira/djinn/backend/internal/application/query/user"
	"github.com/fernandobandeira/djinn/backend/internal/database"
	"log/slog"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the GraphQL resolver root
type Resolver struct {
	// Infrastructure
	DB      PingableDB
	Queries QueriesInterface
	Logger  LoggerInterface

	// Command Handlers
	CreateUserHandler *commandUser.CreateUserHandler
	UpdateUserHandler *commandUser.UpdateUserHandler
	DeleteUserHandler *commandUser.DeleteUserHandler

	// Query Handlers
	GetUserHandler              *queryUser.GetUserHandler
	GetUserByFirebaseUIDHandler *queryUser.GetUserByFirebaseUIDHandler
}

// NewResolver creates a new GraphQL resolver using Wire dependency injection
func NewResolver(
	db *database.DB,
	logger *slog.Logger,
	createUserHandler *commandUser.CreateUserHandler,
	updateUserHandler *commandUser.UpdateUserHandler,
	deleteUserHandler *commandUser.DeleteUserHandler,
	getUserHandler *queryUser.GetUserHandler,
	getUserByFirebaseUIDHandler *queryUser.GetUserByFirebaseUIDHandler,
) *Resolver {
	return &Resolver{
		DB:      db,
		Queries: db.Queries,
		Logger:  logger,

		// Handlers injected by Wire
		CreateUserHandler:           createUserHandler,
		UpdateUserHandler:           updateUserHandler,
		DeleteUserHandler:           deleteUserHandler,
		GetUserHandler:              getUserHandler,
		GetUserByFirebaseUIDHandler: getUserByFirebaseUIDHandler,
	}
}

// NewResolverWithInterfaces creates a resolver with injected interfaces for testing
func NewResolverWithInterfaces(queries QueriesInterface, logger LoggerInterface) *Resolver {
	return &Resolver{
		Queries: queries,
		Logger:  logger,
	}
}