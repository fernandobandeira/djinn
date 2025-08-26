//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

// InitializeApp creates a new application with all dependencies injected
func InitializeApp() (*Application, error) {
	wire.Build(
		ProvideConfig,
		ProvideLogger,
		ProvideDatabase,
		ProvideServer,
		ProvideTracing,
		ProvideApplication,
	)
	return nil, nil
}