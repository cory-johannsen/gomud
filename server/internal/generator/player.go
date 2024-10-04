package generator

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/loader"
)

type PlayerGenerator struct {
	appearanceLoader *loader.AppearanceLoader
	backgroundLoader *loader.BackgroundLoader
}

func NewPlayerGenerator(al *loader.AppearanceLoader, bl *loader.BackgroundLoader) *PlayerGenerator {
	return &PlayerGenerator{
		appearanceLoader: al,
		backgroundLoader: bl,
	}
}

func (g *PlayerGenerator) Generate(name string, pw string, team *domain.Team) (*domain.Player, error) {
	player := domain.NewPlayer(nil, name, pw)
	player.Data[domain.TeamProperty] = team
	player.Data[domain.StatsProperty] = domain.NewStats()
	// generate background
	background, err := g.backgroundLoader.RandomBackground()
	if err != nil {
		return nil, err
	}
	player.Data[domain.BackgroundProperty] = background
	// Generate appearance
	return player, nil
}
