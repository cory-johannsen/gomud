package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Slow struct {
	name        string
	description string
}

func NewSlow() *Slow {
	return &Slow{
		name:        "Slow",
		description: "Whenever a foe is struck by weapons of this Quality, they gain a +10 Base Chance to Dodge or Parry its Damage.",
	}
}

func (e *Slow) Name() string {
	return e.name
}

func (e *Slow) Description() string {
	return e.description
}

func (e *Slow) Applier() domain.Applier {
	return e.Apply
}

func (e *Slow) Apply(state domain.GameState) domain.GameState {
	// - Whenever a foe is struck by weapons of this Quality, they gain a +10 Base Chance to Dodge or Parry its Damage.
	log.Println("applying Slow")
	return state
}

var _ domain.Effect = &Slow{}
