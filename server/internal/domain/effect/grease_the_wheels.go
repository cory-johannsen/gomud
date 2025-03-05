package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type GreasetheWheels struct {
	name        string
	description string
}

func NewGreasetheWheels() *GreasetheWheels {
	return &GreasetheWheels{
		name:        "Grease the Wheels",
		description: "Whenever you fail a Bargain or Counterfeit Test, you may re-roll to generate a better result, but must accept the outcome.",
	}
}

func (e *GreasetheWheels) Name() string {
	return e.name
}

func (e *GreasetheWheels) Description() string {
	return e.description
}

func (e *GreasetheWheels) Applier() domain.Applier {
	return e.Apply
}

func (e *GreasetheWheels) Apply(state domain.GameState) domain.GameState {
	// - Whenever you fail a Bargain or Counterfeit Test, you may re-roll to generate a better result, but must accept the outcome.
	log.Println("applying Grease the Wheels")
	return state
}

var _ domain.Effect = &GreasetheWheels{}
