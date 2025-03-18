//go:build wireinject
// +build wireinject

package main

import (
	eventbus "github.com/asaskevich/EventBus"
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain/effect"
	"github.com/cory-johannsen/gomud/internal/domain/htn"
	"github.com/cory-johannsen/gomud/internal/engine"
	"github.com/cory-johannsen/gomud/internal/event"
	"github.com/cory-johannsen/gomud/internal/generator"
	"github.com/cory-johannsen/gomud/internal/loader"
	"github.com/cory-johannsen/gomud/internal/storage"
	"github.com/google/wire"
)

func InitializeEngine() (*engine.Engine, error) {
	wire.Build(config.NewConfigFromEnv, storage.NewDatabase, storage.NewNPCs, storage.NewPlayers, storage.NewEquipment,
		effect.EffectsSet, effect.NewEffects,
		loader.LoaderSet, loader.NewLoaders,
		generator.NewPlayerGenerator,
		generator.NewDomainGenerator, wire.Bind(new(htn.DomainResolver), new(*generator.DomainGenerator)),
		generator.NewPlannerGenerator, wire.Bind(new(htn.PlannerResolver), new(*generator.PlannerGenerator)),
		eventbus.New, event.NewClock, engine.NewServer, engine.NewEngine)
	return &engine.Engine{}, nil
}
