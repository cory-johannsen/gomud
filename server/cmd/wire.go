//go:build wireinject
// +build wireinject

package main

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/engine"
	"github.com/cory-johannsen/gomud/internal/generator"
	"github.com/cory-johannsen/gomud/internal/loader"
	"github.com/cory-johannsen/gomud/internal/storage"
	"github.com/google/wire"
	goeventbus "github.com/stanipetrosyan/go-eventbus"
)

func InitializeEngine() (*engine.Engine, error) {
	wire.Build(config.NewConfigFromEnv, storage.NewDatabase, storage.NewPlayers,
		loader.NewAppearanceLoader, loader.NewAlignmentLoader, loader.NewArchetypeLoader, loader.NewBackgroundLoader,
		loader.NewInjuryLoader, loader.NewJobLoader, loader.NewRoomLoader, loader.NewSkillLoader, loader.NewTalentLoader,
		loader.NewTeamLoader, loader.NewTraitLoader, loader.NewLoaders,
		generator.NewPlayerGenerator,
		goeventbus.NewEventBus, engine.NewClock, engine.NewServer, engine.NewEngine)
	return &engine.Engine{}, nil
}
