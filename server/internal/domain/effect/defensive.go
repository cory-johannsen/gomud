package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Defensive struct {
	name        string
	description string
}

func NewDefensive() *Defensive {
	return &Defensive{
		name:        "Defensive",
		description: "Shields and weapons of this Quality add a +10 Base Chance to Parry.",
	}
}

func (e *Defensive) Name() string {
	return e.name
}

func (e *Defensive) Description() string {
	return e.description
}

func (e *Defensive) Applier() domain.Applier {
	return e.Apply
}

func (e *Defensive) Apply(state domain.GameState) domain.GameState {
	// - Shields and weapons of this Quality add a +10 Base Chance to Parry.
	log.Println("applying Defensive")
	return state
}

var _ domain.Effect = &Defensive{}
