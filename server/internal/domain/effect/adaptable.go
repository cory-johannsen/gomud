package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Adaptable struct {
	name        string
	description string
}

func NewAdaptable() *Adaptable {
	return &Adaptable{
		name:        "Adaptable",
		description: "Whenever weapons of this Quality are held in two hands instead of one for Attack Actions, add +1 to Damage.",
	}
}

func (e *Adaptable) Name() string {
	return e.name
}

func (e *Adaptable) Description() string {
	return e.description
}

func (e *Adaptable) Applier() domain.Applier {
	return e.Apply
}

func (e *Adaptable) Apply(state domain.GameState) domain.GameState {
	// - Whenever weapons of this Quality are held in two hands instead of one for Attack Actions, add +1 to Damage.
	log.Println("applying Adaptable")
	return state
}

var _ domain.Effect = &Adaptable{}
