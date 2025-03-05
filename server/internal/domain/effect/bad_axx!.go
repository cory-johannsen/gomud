package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type BadAxx struct {
	name        string
	description string
}

func NewBadAxx() *BadAxx {
	return &BadAxx{
		name:        "Bad Axx!",
		description: "When you hold a one-handed melee weapon in either hand and fail a Martial Melee or Simple Melee Test, you may re-roll to generate a better result, but must accept the outcome. However, if you do not possess Ambidexterity, you must flip the results to fail the re-rolled Skill Test.",
	}
}

func (e *BadAxx) Name() string {
	return e.name
}

func (e *BadAxx) Description() string {
	return e.description
}

func (e *BadAxx) Applier() domain.Applier {
	return e.Apply
}

func (e *BadAxx) Apply(state domain.GameState) domain.GameState {
	// - When you hold a one-handed melee weapon in either hand and fail a Martial Melee or Simple Melee Test, you may re-roll to generate a better result, but must accept the outcome. However, if you do not possess Ambidexterity, you must flip the results to fail the re-rolled Skill Test.
	log.Println("applying Bad Axx!")
	return state
}

var _ domain.Effect = &BadAxx{}
