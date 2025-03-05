package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Vicious struct {
	name        string
	description string
}

func NewVicious() *Vicious {
	return &Vicious{
		name:        "Vicious",
		description: "Weapons of this Quality grant an additional 1D6 Chaos Die to determine whether you inflict an Injury upon a foe.",
	}
}

func (e *Vicious) Name() string {
	return e.name
}

func (e *Vicious) Description() string {
	return e.description
}

func (e *Vicious) Applier() domain.Applier {
	return e.Apply
}

func (e *Vicious) Apply(state domain.GameState) domain.GameState {
	// - Weapons of this Quality grant an additional 1D6 Chaos Die to determine whether you inflict an Injury upon a foe.
	log.Println("applying Vicious")
	return state
}

var _ domain.Effect = &Vicious{}
