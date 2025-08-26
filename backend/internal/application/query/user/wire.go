//go:build !wireinject
// +build !wireinject

package user

import "github.com/google/wire"

// QuerySet provides all user query handlers
var QuerySet = wire.NewSet(
	NewGetUserHandler,
	NewGetUserByFirebaseUIDHandler,
)