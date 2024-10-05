package generator

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/loader"
	"log"
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
	log.Printf("generating player %s", name)
	player := domain.NewPlayer(nil, name, pw, nil)
	player.Data[domain.TeamProperty] = team
	player.Data[domain.StatsProperty] = domain.NewStats()
	// generate background
	background, err := g.backgroundLoader.RandomBackground()
	if err != nil {
		log.Printf("failed to generate background for player %s: %s", name, err)
		return nil, err
	}
	player.Data[domain.BackgroundProperty] = background
	// Generate appearance
	return player, nil
}
