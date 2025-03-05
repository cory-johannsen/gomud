package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Cursed struct {
	name        string
	description string
}

func NewCursed() *Cursed {
	return &Cursed{
		name:        "Cursed",
		description: "Whenever you intend to sacrifice a Fortune Point, roll a 1D6 Chaos Die. If the result is a face ‘6’, you must use two Fortune Points instead of one.",
	}
}

func (e *Cursed) Name() string {
	return e.name
}

func (e *Cursed) Description() string {
	return e.description
}

func (e *Cursed) Applier() domain.Applier {
	return e.Apply
}

func (e *Cursed) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
	log.Println("applying Cursed")
	return state
}

var _ domain.Effect = &Cursed{}
