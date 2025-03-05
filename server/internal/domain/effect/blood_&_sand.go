package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type BloodSand struct {
	name        string
	description string
}

func NewBloodSand() *BloodSand {
	return &BloodSand{
		name:        "Blood & Sand",
		description: "Whenever you spend a Fortune Point, you move one step up the Damage & Peril Condition Tracks positively.",
	}
}

func (e *BloodSand) Name() string {
	return e.name
}

func (e *BloodSand) Description() string {
	return e.description
}

func (e *BloodSand) Applier() domain.Applier {
	return e.Apply
}

func (e *BloodSand) Apply(state domain.GameState) domain.GameState {
	// - Whenever you spend a Fortune Point, you move one step up the Damage & Peril Condition Tracks positively.
	log.Println("applying Blood & Sand")
	return state
}

var _ domain.Effect = &BloodSand{}
