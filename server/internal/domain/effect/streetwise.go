package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Streetwise struct {
	name        string
	description string
}

func NewStreetwise() *Streetwise {
	return &Streetwise{
		name:        "Streetwise",
		description: "When you attempt to hide in urban environments, you gain a +20 Base Chance to Stealth Tests.",
	}
}

func (e *Streetwise) Name() string {
	return e.name
}

func (e *Streetwise) Description() string {
	return e.description
}

func (e *Streetwise) Applier() domain.Applier {
	return e.Apply
}

func (e *Streetwise) Apply(state domain.GameState) domain.GameState {
	// - When you attempt to hide in urban environments, you gain a +20 Base Chance to Stealth Tests.
	log.Println("applying Streetwise")
	return state
}

var _ domain.Effect = &Streetwise{}
