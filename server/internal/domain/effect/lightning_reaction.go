package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type LightningReaction struct {
	name        string
	description string
}

func NewLightningReaction() *LightningReaction {
	return &LightningReaction{
		name:        "Lightning Reaction",
		description: "In combat, you gain 1 additional Action Point. However, it can only be used to Dodge and Parry. This Action Point refreshes at the beginning of your Turn.",
	}
}

func (e *LightningReaction) Name() string {
	return e.name
}

func (e *LightningReaction) Description() string {
	return e.description
}

func (e *LightningReaction) Applier() domain.Applier {
	return e.Apply
}

func (e *LightningReaction) Apply(state domain.GameState) domain.GameState {
	// - In combat, you gain 1 additional Action Point. However, it can only be used to Dodge and Parry. This Action Point refreshes at the beginning of your Turn.
	log.Println("applying Lightning Reaction")
	return state
}

var _ domain.Effect = &LightningReaction{}
