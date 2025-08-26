//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	
	"github.com/fernandobandeira/djinn/backend/internal/adapter/repository/postgres"
	"github.com/fernandobandeira/djinn/backend/internal/application/command/user"
	queryUser "github.com/fernandobandeira/djinn/backend/internal/application/query/user"
	domainUser "github.com/fernandobandeira/djinn/backend/internal/domain/user"
	"github.com/fernandobandeira/djinn/backend/internal/graph/resolver"
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/config"
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/logging"
)

// InfrastructureProviderSet groups infrastructure providers
var InfrastructureProviderSet = wire.NewSet(
	config.Load,
	logging.NewLogger,
	ProvideDatabase,
	ProvideQueries,
	ProvideMonitoring,
)

// RepositoryProviderSet groups repository providers
var RepositoryProviderSet = wire.NewSet(
	postgres.NewUserRepository,
	wire.Bind(new(domainUser.Repository), new(*postgres.UserRepository)),
)

// DomainProviderSet groups domain providers
var DomainProviderSet = wire.NewSet(
	domainUser.NewService,
)

// ApplicationProviderSet groups application providers
var ApplicationProviderSet = wire.NewSet(
	// Command handlers
	user.NewCreateUserHandler,
	user.NewUpdateUserHandler,
	user.NewDeleteUserHandler,
	
	// Query handlers  
	queryUser.NewGetUserHandler,
	queryUser.NewGetUserByFirebaseUIDHandler,
)

// PresentationProviderSet groups presentation providers
var PresentationProviderSet = wire.NewSet(
	resolver.NewResolver,
	ProvideServer,
)

// InitializeApp creates a new application with all dependencies injected
func InitializeApp() (*Application, func(), error) {
	wire.Build(
		// Layer-based provider sets
		InfrastructureProviderSet,
		RepositoryProviderSet,
		DomainProviderSet,
		ApplicationProviderSet,
		PresentationProviderSet,
		
		// Application assembly
		ProvideApplication,
	)
	return nil, nil, nil
}