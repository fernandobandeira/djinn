//go:build !wireinject
// +build !wireinject

package resolver

import (
	"github.com/google/wire"
	
	"github.com/fernandobandeira/djinn/backend/internal/adapter/repository/postgres"
	"github.com/fernandobandeira/djinn/backend/internal/application"
	domainUser "github.com/fernandobandeira/djinn/backend/internal/domain/user"
)

// GraphQLSet provides all GraphQL resolver dependencies
var GraphQLSet = wire.NewSet(
	NewResolver,
	
	// Include all the domain, application and adapter sets
	domainUser.DomainSet,
	application.ApplicationSet,
	postgres.RepositorySet,
)