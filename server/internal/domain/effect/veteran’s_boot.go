package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type VeteransBoot struct {
	name        string
	description string
}

func NewVeteransBoot() *VeteransBoot {
	return &VeteransBoot{
		name:        "Veteran’s Boot",
		description: "You cannot Charge, Maneuver or Run with Movement Actions without spending an additional Action Point.",
	}
}

func (e *VeteransBoot) Name() string {
	return e.name
}

func (e *VeteransBoot) Description() string {
	return e.description
}

func (e *VeteransBoot) Applier() domain.Applier {
	return e.Apply
}

func (e *VeteransBoot) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
	log.Println("applying Veteran’s Boot")
	return state
}

var _ domain.Effect = &VeteransBoot{}
