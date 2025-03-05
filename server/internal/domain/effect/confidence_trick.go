package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type ConfidenceTrick struct {
	name        string
	description string
}

func NewConfidenceTrick() *ConfidenceTrick {
	return &ConfidenceTrick{
		name:        "Confidence Trick",
		description: "You may flip the results to succeed at Guile Tests. When you succeed, it is always considered a Critical Success. Furthermore, you can influence a number of people with the Guile Skill equal to your [FB] times three – this includes using Dirty Tricks during combat.",
	}
}

func (e *ConfidenceTrick) Name() string {
	return e.name
}

func (e *ConfidenceTrick) Description() string {
	return e.description
}

func (e *ConfidenceTrick) Applier() domain.Applier {
	return e.Apply
}

func (e *ConfidenceTrick) Apply(state domain.GameState) domain.GameState {
	// - You may flip the results to succeed at Guile Tests. When you succeed, it is always considered a Critical Success. Furthermore, you can influence a number of people with the Guile Skill equal to your [FB] times three – this includes using Dirty Tricks during combat.
	log.Println("applying Confidence Trick")
	return state
}

var _ domain.Effect = &ConfidenceTrick{}
