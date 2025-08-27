//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/config"
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/logging"
)

// InitializeApp creates a new lifecycle-managed application
func InitializeApp() (*Application, func(), error) {
	wire.Build(
		// Basic infrastructure
		config.Load,
		logging.NewLogger,
		
		// Application with lifecycle management
		NewApplication,
	)
	return nil, nil, nil
}