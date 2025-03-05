package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Weak struct {
	name        string
	description string
}

func NewWeak() *Weak {
	return &Weak{
		name:        "Weak",
		description: "Weapons of this Quality can only inflict Moderate or Serious Injuries, never Grievous Injuries.",
	}
}

func (e *Weak) Name() string {
	return e.name
}

func (e *Weak) Description() string {
	return e.description
}

func (e *Weak) Applier() domain.Applier {
	return e.Apply
}

func (e *Weak) Apply(state domain.GameState) domain.GameState {
	// - Weapons of this Quality can only inflict Moderate or Serious Injuries, never Grievous Injuries.
	log.Println("applying Weak")
	return state
}

var _ domain.Effect = &Weak{}
