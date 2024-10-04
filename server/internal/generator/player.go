package generator

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/loader"
)

type PlayerGenerator struct {
	appearanceLoader *loader.AppearanceLoader
}

func NewPlayerGenerator(al *loader.AppearanceLoader) *PlayerGenerator {
	return &PlayerGenerator{
		appearanceLoader: al,
	}
}

func (pg *PlayerGenerator) Generate(name string, pw string, team *domain.Team) (*domain.Player, error) {
	player := domain.NewPlayer(nil, name, pw)
	player.Data["team"] = team.Name

	return player, nil
}
