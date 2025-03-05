package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type ThickLining struct {
	name        string
	description string
}

func NewThickLining() *ThickLining {
	return &ThickLining{
		name:        "Thick Lining",
		description: "You are immune to Poisons of all types.,",
	}
}

func (e *ThickLining) Name() string {
	return e.name
}

func (e *ThickLining) Description() string {
	return e.description
}

func (e *ThickLining) Applier() domain.Applier {
	return e.Apply
}

func (e *ThickLining) Apply(state domain.GameState) domain.GameState {
	// - You are immune to Poisons of all types.,
	log.Println("applying Thick Lining")
	return state
}

var _ domain.Effect = &ThickLining{}
