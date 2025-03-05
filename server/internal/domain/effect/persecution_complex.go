package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type PersecutionComplex struct {
	name        string
	description string
}

func NewPersecutionComplex() *PersecutionComplex {
	return &PersecutionComplex{
		name:        "Persecution Complex",
		description: "You cannot rest to recover from Peril in urban environments, unless you take a dose of laudanum before resting.",
	}
}

func (e *PersecutionComplex) Name() string {
	return e.name
}

func (e *PersecutionComplex) Description() string {
	return e.description
}

func (e *PersecutionComplex) Applier() domain.Applier {
	return e.Apply
}

func (e *PersecutionComplex) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
	log.Println("applying Persecution Complex")
	return state
}

var _ domain.Effect = &PersecutionComplex{}
