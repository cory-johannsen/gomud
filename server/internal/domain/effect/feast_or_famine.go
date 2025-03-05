package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type FeastorFamine struct {
	name        string
	description string
}

func NewFeastorFamine() *FeastorFamine {
	return &FeastorFamine{
		name:        "Feast or Famine",
		description: "You can attempt a Survival or Tradecraft Test to cook for others. If successful, a number of allies equal to your [FB] move one step up the Damage Condition Track positively, and recover their Peril Condition Track to Unhindered, whether they’re resting in a dangerous place or not. If your roll is a Critical Success, you affect a number of allies equal to three times your [FB] instead. A Character may only gain the benefit of this meal once per day, which takes an hour of time to prepare.",
	}
}

func (e *FeastorFamine) Name() string {
	return e.name
}

func (e *FeastorFamine) Description() string {
	return e.description
}

func (e *FeastorFamine) Applier() domain.Applier {
	return e.Apply
}

func (e *FeastorFamine) Apply(state domain.GameState) domain.GameState {
	// - You can attempt a Survival or Tradecraft Test to cook for others. If successful, a number of allies equal to your [FB] move one step up the Damage Condition Track positively, and recover their Peril Condition Track to Unhindered, whether they’re resting in a dangerous place or not. If your roll is a Critical Success, you affect a number of allies equal to three times your [FB] instead. A Character may only gain the benefit of this meal once per day, which takes an hour of time to prepare.
	log.Println("applying Feast or Famine")
	return state
}

var _ domain.Effect = &FeastorFamine{}
