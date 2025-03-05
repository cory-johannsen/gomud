package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Shrapnel struct {
	name        string
	description string
}

func NewShrapnel() *Shrapnel {
	return &Shrapnel{
		name:        "Shrapnel",
		description: "Whenever a weapon with this Quality is fired, it affects multiple targets in a Cone Template.",
	}
}

func (e *Shrapnel) Name() string {
	return e.name
}

func (e *Shrapnel) Description() string {
	return e.description
}

func (e *Shrapnel) Applier() domain.Applier {
	return e.Apply
}

func (e *Shrapnel) Apply(state domain.GameState) domain.GameState {
	// - Whenever a weapon with this Quality is fired, it affects multiple targets in a Cone Template.
	log.Println("applying Shrapnel")
	return state
}

var _ domain.Effect = &Shrapnel{}
