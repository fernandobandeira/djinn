//go:build !wireinject
// +build !wireinject

package application

import (
	"github.com/google/wire"
	
	commandUser "github.com/fernandobandeira/djinn/backend/internal/application/command/user"
	queryUser "github.com/fernandobandeira/djinn/backend/internal/application/query/user"
)

// ApplicationSet provides all application layer services
var ApplicationSet = wire.NewSet(
	// Command handlers
	commandUser.CommandSet,
	
	// Query handlers
	queryUser.QuerySet,
)