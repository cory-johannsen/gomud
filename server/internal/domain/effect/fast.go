package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Fast struct {
	name        string
	description string
}

func NewFast() *Fast {
	return &Fast{
		name:        "Fast",
		description: "Whenever a foe is struck by weapons of this Quality, they suffer a -10 Base Chance to Dodge or Parry.",
	}
}

func (e *Fast) Name() string {
	return e.name
}

func (e *Fast) Description() string {
	return e.description
}

func (e *Fast) Applier() domain.Applier {
	return e.Apply
}

func (e *Fast) Apply(state domain.GameState) domain.GameState {
	// - Whenever a foe is struck by weapons of this Quality, they suffer a -10 Base Chance to Dodge or Parry.
	log.Println("applying Fast")
	return state
}

var _ domain.Effect = &Fast{}
