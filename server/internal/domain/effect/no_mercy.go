package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type NoMercy struct {
	name        string
	description string
}

func NewNoMercy() *NoMercy {
	return &NoMercy{
		name:        "No Mercy",
		description: "When you Injure a foe with a melee weapon, you inflict two Injuries instead of one.",
	}
}

func (e *NoMercy) Name() string {
	return e.name
}

func (e *NoMercy) Description() string {
	return e.description
}

func (e *NoMercy) Applier() domain.Applier {
	return e.Apply
}

func (e *NoMercy) Apply(state domain.GameState) domain.GameState {
	// - When you Injure a foe with a melee weapon, you inflict two Injuries instead of one.
	log.Println("applying No Mercy")
	return state
}

var _ domain.Effect = &NoMercy{}
