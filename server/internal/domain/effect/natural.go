package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Natural struct {
	name        string
	description string
}

func NewNatural() *Natural {
	return &Natural{
		name:        "Natural",
		description: "Armor of this Quality adds a +10 Base Chance to Dodge attacks.",
	}
}

func (e *Natural) Name() string {
	return e.name
}

func (e *Natural) Description() string {
	return e.description
}

func (e *Natural) Applier() domain.Applier {
	return e.Apply
}

func (e *Natural) Apply(state domain.GameState) domain.GameState {
	// - Armor of this Quality adds a +10 Base Chance to Dodge attacks.
	log.Println("applying Natural")
	return state
}

var _ domain.Effect = &Natural{}
