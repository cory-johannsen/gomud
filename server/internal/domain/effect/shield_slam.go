package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type ShieldSlam struct {
	name        string
	description string
}

func NewShieldSlam() *ShieldSlam {
	return &ShieldSlam{
		name:        "Shield Slam",
		description: "When you wield a shield, the melee weapon you wield in your primary hand gains the Powerful Quality.",
	}
}

func (e *ShieldSlam) Name() string {
	return e.name
}

func (e *ShieldSlam) Description() string {
	return e.description
}

func (e *ShieldSlam) Applier() domain.Applier {
	return e.Apply
}

func (e *ShieldSlam) Apply(state domain.GameState) domain.GameState {
	// - When you wield a shield, the melee weapon you wield in your primary hand gains the Powerful Quality.
	log.Println("applying Shield Slam")
	return state
}

var _ domain.Effect = &ShieldSlam{}
