//go:build wireinject
// +build wireinject

package main

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/engine"
	"github.com/cory-johannsen/gomud/internal/storage"
	"github.com/google/wire"
)

func InitializeEngine() (*engine.Engine, error) {
	wire.Build(config.NewConfigFromEnv, storage.NewDatabase, storage.NewPlayers,
		engine.NewServer, engine.NewEngine)
	return &engine.Engine{}, nil
}
