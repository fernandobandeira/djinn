//go:build !wireinject
// +build !wireinject

package postgres

import (
	"github.com/google/wire"
	
	domainUser "github.com/fernandobandeira/djinn/backend/internal/domain/user"
)

// RepositorySet provides all PostgreSQL repository implementations
var RepositorySet = wire.NewSet(
	NewUserRepository,
	
	// Bind to domain interfaces
	wire.Bind(new(domainUser.Repository), new(*UserRepository)),
)