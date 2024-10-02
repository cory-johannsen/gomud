package generator

import "github.com/cory-johannsen/gomud/internal/loader"

type PlayerGenerator struct {
	appearanceLoader *loader.AppearanceLoader
}

func NewPlayerGenerator(al *loader.AppearanceLoader) *PlayerGenerator {
	return &PlayerGenerator{
		appearanceLoader: al,
	}
}
