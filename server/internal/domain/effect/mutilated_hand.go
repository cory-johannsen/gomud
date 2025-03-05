package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type MutilatedHand struct {
	name        string
	description string
}

func NewMutilatedHand() *MutilatedHand {
	return &MutilatedHand{
		name:        "Mutilated Hand",
		description: "Until fully Recuperated, you cannot use your primary hand as you’re in pain. You must undergo a successful surgery or suffer the consequences. Once a Mutilated Hand has undergone a failed surgery, you lose 1D6-1 fingers. For every finger lost, you permanently lose 1% Agility. If you lose all fingers, you gain the Veteran’s Hand Drawback. If you already have this Drawback, you permanently lose 9% Agility.",
	}
}

func (e *MutilatedHand) Name() string {
	return e.name
}

func (e *MutilatedHand) Description() string {
	return e.description
}

func (e *MutilatedHand) Applier() domain.Applier {
	return e.Apply
}

func (e *MutilatedHand) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you cannot use your primary hand as you’re in pain. You must undergo a successful surgery or suffer the consequences. Once a Mutilated Hand has undergone a failed surgery, you lose 1D6-1 fingers. For every finger lost, you permanently lose 1% Agility. If you lose all fingers, you gain the Veteran’s Hand Drawback. If you already have this Drawback, you permanently lose 9% Agility.
	log.Println("applying Mutilated Hand")
	return state
}

var _ domain.Effect = &MutilatedHand{}
