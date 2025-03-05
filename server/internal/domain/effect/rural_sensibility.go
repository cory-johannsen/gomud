package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type RuralSensibility struct {
	name        string
	description string
}

func NewRuralSensibility() *RuralSensibility {
	return &RuralSensibility{
		name:        "Rural Sensibility",
		description: "When you attempt to hide in rural environments, you gain a +20 Base Chance to Stealth Tests.",
	}
}

func (e *RuralSensibility) Name() string {
	return e.name
}

func (e *RuralSensibility) Description() string {
	return e.description
}

func (e *RuralSensibility) Applier() domain.Applier {
	return e.Apply
}

func (e *RuralSensibility) Apply(state domain.GameState) domain.GameState {
	// - When you attempt to hide in rural environments, you gain a +20 Base Chance to Stealth Tests.
	log.Println("applying Rural Sensibility")
	return state
}

var _ domain.Effect = &RuralSensibility{}
