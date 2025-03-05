package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type DetachedEar struct {
	name        string
	description string
}

func NewDetachedEar() *DetachedEar {
	return &DetachedEar{
		name:        "Detached Ear",
		description: "Until fully Recuperated, you cannot hear as you’re in pain. You must undergo a successful surgery or suffer the consequences. Once a Detached Ear has undergone a failed surgery, you suffer the Crop Ear Drawback. If you already have this Drawback, you permanently lose 9% Perception.",
	}
}

func (e *DetachedEar) Name() string {
	return e.name
}

func (e *DetachedEar) Description() string {
	return e.description
}

func (e *DetachedEar) Applier() domain.Applier {
	return e.Apply
}

func (e *DetachedEar) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you cannot hear as you’re in pain. You must undergo a successful surgery or suffer the consequences. Once a Detached Ear has undergone a failed surgery, you suffer the Crop Ear Drawback. If you already have this Drawback, you permanently lose 9% Perception.
	log.Println("applying Detached Ear")
	return state
}

var _ domain.Effect = &DetachedEar{}
