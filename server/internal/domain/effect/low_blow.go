package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type LowBlow struct {
	name        string
	description string
}

func NewLowBlow() *LowBlow {
	return &LowBlow{
		name:        "Low Blow",
		description: "Whenever you make a successful Melee Attack, you force a foe to Resist a Stunning Blow. You must be Engaged with a foe in order to use a Low Blow.,",
	}
}

func (e *LowBlow) Name() string {
	return e.name
}

func (e *LowBlow) Description() string {
	return e.description
}

func (e *LowBlow) Applier() domain.Applier {
	return e.Apply
}

func (e *LowBlow) Apply(state domain.GameState) domain.GameState {
	// - Whenever you make a successful Melee Attack, you force a foe to Resist a Stunning Blow. You must be Engaged with a foe in order to use a Low Blow.,
	log.Println("applying Low Blow")
	return state
}

var _ domain.Effect = &LowBlow{}
