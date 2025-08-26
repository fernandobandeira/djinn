//go:build !wireinject
// +build !wireinject

package user

import "github.com/google/wire"

// CommandSet provides all user command handlers
var CommandSet = wire.NewSet(
	NewCreateUserHandler,
	NewUpdateUserHandler,
	NewDeleteUserHandler,
)