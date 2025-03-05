package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type MaimedFoot struct {
	name        string
	description string
}

func NewMaimedFoot() *MaimedFoot {
	return &MaimedFoot{
		name:        "Maimed Foot",
		description: "Until fully Recuperated, you cannot use any weapon with an Encumbrance Value of 2 or more, as you’re in pain. You must undergo a successful surgery or suffer the consequences. Once a Maimed Foot has undergone a failed surgery, you lose 1D6-1 toes. For every toe lost, you permanently lose 1% Brawn. If you lose all toes, you gain the Veteran’s Boot Drawback. If you already have this Drawback, you permanently lose 9% Brawn.",
	}
}

func (e *MaimedFoot) Name() string {
	return e.name
}

func (e *MaimedFoot) Description() string {
	return e.description
}

func (e *MaimedFoot) Applier() domain.Applier {
	return e.Apply
}

func (e *MaimedFoot) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you cannot use any weapon with an Encumbrance Value of 2 or more, as you’re in pain. You must undergo a successful surgery or suffer the consequences. Once a Maimed Foot has undergone a failed surgery, you lose 1D6-1 toes. For every toe lost, you permanently lose 1% Brawn. If you lose all toes, you gain the Veteran’s Boot Drawback. If you already have this Drawback, you permanently lose 9% Brawn.
	log.Println("applying Maimed Foot")
	return state
}

var _ domain.Effect = &MaimedFoot{}
