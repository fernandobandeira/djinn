//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/fernandobandeira/djinn/backend/internal/config"
	"github.com/fernandobandeira/djinn/backend/internal/database"
	"github.com/fernandobandeira/djinn/backend/internal/logger"
	"github.com/fernandobandeira/djinn/backend/internal/server"
)

// InitializeApp creates a new server with all dependencies injected
func InitializeApp() (*Application, error) {
	wire.Build(
		config.Load,
		logger.New,
		provideDatabase,
		server.NewServer,
		NewApplication,
	)
	return nil, nil
}

// provideDatabase creates a database connection from config
func provideDatabase(cfg *config.Config) (*database.DB, error) {
	return database.Connect(cfg.PgBouncerURL)
}