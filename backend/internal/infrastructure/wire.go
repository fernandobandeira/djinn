//go:build !wireinject
// +build !wireinject

package infrastructure

import (
	"github.com/google/wire"
	
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/config"
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/logging"
)

// InfrastructureSet provides all infrastructure components
var InfrastructureSet = wire.NewSet(
	config.Load,  // Configuration loading
	logging.NewLogger,  // Logger
)