// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/engine"
	"github.com/cory-johannsen/gomud/internal/generator"
	"github.com/cory-johannsen/gomud/internal/loader"
	"github.com/cory-johannsen/gomud/internal/storage"
)

// Injectors from wire.go:

func InitializeEngine() (*engine.Engine, error) {
	configConfig, err := config.NewConfigFromEnv()
	if err != nil {
		return nil, err
	}
	database, err := storage.NewDatabase(configConfig)
	if err != nil {
		return nil, err
	}
	players := storage.NewPlayers(database)
	appearanceLoader := loader.NewAppearanceLoader(configConfig)
	playerGenerator := generator.NewPlayerGenerator(appearanceLoader)
	server := engine.NewServer(configConfig, database, players, appearanceLoader, playerGenerator)
	engineEngine := engine.NewEngine(configConfig, server)
	return engineEngine, nil
}
