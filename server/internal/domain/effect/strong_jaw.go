package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type StrongJaw struct {
	name        string
	description string
}

func NewStrongJaw() *StrongJaw {
	return &StrongJaw{
		name:        "Strong Jaw",
		description: "When attempting to Resist Perilous Stunts, you gain a +20 Base Chance to succeed.",
	}
}

func (e *StrongJaw) Name() string {
	return e.name
}

func (e *StrongJaw) Description() string {
	return e.description
}

func (e *StrongJaw) Applier() domain.Applier {
	return e.Apply
}

func (e *StrongJaw) Apply(state domain.GameState) domain.GameState {
	// - When attempting to Resist Perilous Stunts, you gain a +20 Base Chance to succeed.
	log.Println("applying Strong Jaw")
	return state
}

var _ domain.Effect = &StrongJaw{}
