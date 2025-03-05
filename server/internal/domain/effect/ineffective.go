package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Ineffective struct {
	name        string
	description string
}

func NewIneffective() *Ineffective {
	return &Ineffective{
		name:        "Ineffective",
		description: "Weapons of this Quality cannot deal Damage or inflict Injuries.",
	}
}

func (e *Ineffective) Name() string {
	return e.name
}

func (e *Ineffective) Description() string {
	return e.description
}

func (e *Ineffective) Applier() domain.Applier {
	return e.Apply
}

func (e *Ineffective) Apply(state domain.GameState) domain.GameState {
	// - Weapons of this Quality cannot deal Damage or inflict Injuries.
	log.Println("applying Ineffective")
	return state
}

var _ domain.Effect = &Ineffective{}
