package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type HigherMysteries struct {
	name        string
	description string
}

func NewHigherMysteries() *HigherMysteries {
	return &HigherMysteries{
		name:        "Higher Mysteries",
		description: "While Channeling Power, you can voluntarily remove any number of Chaos Dice. However, you suffer 3 additional Corruption for each Chaos Die you remove.",
	}
}

func (e *HigherMysteries) Name() string {
	return e.name
}

func (e *HigherMysteries) Description() string {
	return e.description
}

func (e *HigherMysteries) Applier() domain.Applier {
	return e.Apply
}

func (e *HigherMysteries) Apply(state domain.GameState) domain.GameState {
	// - While Channeling Power, you can voluntarily remove any number of Chaos Dice. However, you suffer 3 additional Corruption for each Chaos Die you remove.
	log.Println("applying Higher Mysteries")
	return state
}

var _ domain.Effect = &HigherMysteries{}
