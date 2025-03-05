package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Throwing struct {
	name        string
	description string
}

func NewThrowing() *Throwing {
	return &Throwing{
		name:        "Throwing",
		description: "Weapons of this Quality do not have a Medium or Long Distance increment for ranged weapons.",
	}
}

func (e *Throwing) Name() string {
	return e.name
}

func (e *Throwing) Description() string {
	return e.description
}

func (e *Throwing) Applier() domain.Applier {
	return e.Apply
}

func (e *Throwing) Apply(state domain.GameState) domain.GameState {
	// - Weapons of this Quality do not have a Medium or Long Distance increment for ranged weapons.
	log.Println("applying Throwing")
	return state
}

var _ domain.Effect = &Throwing{}
