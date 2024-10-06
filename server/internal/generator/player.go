package generator

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/loader"
	"log"
	"math/rand"
)

type PlayerGenerator struct {
	loaders *loader.Loaders
}

func NewPlayerGenerator(loaders *loader.Loaders) *PlayerGenerator {
	return &PlayerGenerator{
		loaders: loaders,
	}
}

func (g *PlayerGenerator) Generate(name string, pw string, team *domain.Team, takeDrawback bool) (*domain.Player, error) {
	log.Printf("generating player %s", name)
	player := domain.NewPlayer(nil, name, pw, nil)
	player.Data[domain.TeamProperty] = team
	player.Data[domain.StatsProperty] = domain.NewStats()
	// generate background
	background, err := g.loaders.BackgroundLoader.RandomBackground()
	if err != nil {
		log.Printf("failed to generate background for player %s: %s", name, err)
		return nil, err
	}
	player.Data[domain.BackgroundProperty] = background
	// generate birth season
	season := domain.RandomSeason()
	player.Data[domain.BirthSeasonProperty] = season

	age := rand.Intn(80) + 18
	player.Data[domain.AgeProperty] = &domain.BaseProperty{Val: age}

	// generate appearance
	marks, err := g.loaders.AppearanceLoader.LoadDistinguishingMarks()
	if err != nil {
		log.Printf("failed to load distinguishing marks: %s", err)
		return nil, err
	}
	player.Data[domain.DistinguishingMarkProperty] = marks.Random(age)

	tats, err := g.loaders.AppearanceLoader.LoadTattoos()
	if err != nil {
		log.Printf("failed to load tattoos: %s", err)
		return nil, err
	}
	tat := tats[season].Random()
	player.Data[domain.TattooProperty] = &tat

	if takeDrawback {
		drawbacks, err := g.loaders.AppearanceLoader.LoadDrawbacks()
		if err != nil {
			log.Printf("failed to load drawbacks: %s", err)
			return nil, err
		}
		player.Data[domain.DrawbackProperty] = drawbacks.Random()
	}

	archetypes, err := g.loaders.ArchetypeLoader.LoadArchetypes()
	if err != nil {
		log.Printf("failed to load archetypes: %s", err)
		return nil, err
	}
	archetype := archetypes.Random()
	player.Data[domain.ArchetypeProperty] = archetype

	job, err := g.loaders.JobLoader.RandomJob(archetype)
	if err != nil {
		log.Printf("failed to load jobs: %s", err)
		return nil, err
	}
	player.Data[domain.JobProperty] = job

	return player, nil
}
