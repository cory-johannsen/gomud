package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Instincts struct {
	name        string
	description string
}

func NewInstincts() *Instincts {
	return &Instincts{
		name:        "Instincts",
		description: "While in fog, mist or smoke, you ignore the normal penalties associated with impaired vision.",
	}
}

func (e *Instincts) Name() string {
	return e.name
}

func (e *Instincts) Description() string {
	return e.description
}

func (e *Instincts) Applier() domain.Applier {
	return e.Apply
}

func (e *Instincts) Apply(state domain.GameState) domain.GameState {
	// - While in fog, mist or smoke, you ignore the normal penalties associated with impaired vision.
	log.Println("applying Instincts")
	return state
}

var _ domain.Effect = &Instincts{}
