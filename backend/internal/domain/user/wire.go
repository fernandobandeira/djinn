//go:build !wireinject
// +build !wireinject

package user

import (
	"github.com/google/wire"
)

// DomainSet provides all user domain services
var DomainSet = wire.NewSet(
	NewService,
)