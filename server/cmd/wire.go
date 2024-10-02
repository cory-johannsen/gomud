//go:build wireinject
// +build wireinject

package main

import (
	"github.com/cory-johannsen/gomud/internal/engine"
	"github.com/google/wire"
)

func InitializeEngine() (*engine.Engine, error) {
	wire.Build(engine.NewConfigFromEnv, engine.NewDatabase, engine.NewServer, engine.NewEngine)
	return &engine.Engine{}, nil
}
