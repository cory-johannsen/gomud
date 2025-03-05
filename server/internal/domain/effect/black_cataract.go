package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type BlackCataract struct {
	name        string
	description string
}

func NewBlackCataract() *BlackCataract {
	return &BlackCataract{
		name:        "Black Cataract",
		description: "Whenever you miss with Attack Actions using a ranged weapon, you must re-roll the result with the same Difficulty Rating. If it is a success, you strike a random ally who is Engaged with your target.",
	}
}

func (e *BlackCataract) Name() string {
	return e.name
}

func (e *BlackCataract) Description() string {
	return e.description
}

func (e *BlackCataract) Applier() domain.Applier {
	return e.Apply
}

func (e *BlackCataract) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
	log.Println("applying Black Cataract")
	return state
}

var _ domain.Effect = &BlackCataract{}
