//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	
	"github.com/fernandobandeira/djinn/backend/internal/adapter/repository/postgres"
	"github.com/fernandobandeira/djinn/backend/internal/application"
	domainUser "github.com/fernandobandeira/djinn/backend/internal/domain/user"
	"github.com/fernandobandeira/djinn/backend/internal/graph/resolver"
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure"
)

// InitializeApp creates a new application with all dependencies injected
func InitializeApp() (*Application, error) {
	wire.Build(
		// Infrastructure
		infrastructure.InfrastructureSet,
		ProvideDatabase,
		
		// Repository
		postgres.RepositorySet,
		
		// Domain
		domainUser.DomainSet,
		
		// Application
		application.ApplicationSet,
		
		// GraphQL
		resolver.NewResolver,
		
		// Server
		ProvideServer,
		ProvideTracing,
		ProvideApplication,
	)
	return nil, nil
}