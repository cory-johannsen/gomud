package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type RabbleRousing struct {
	name        string
	description string
}

func NewRabbleRousing() *RabbleRousing {
	return &RabbleRousing{
		name:        "Rabble-Rousing",
		description: "Whenever you fail a Guile or Leadership Test, you may re-roll to generate a better result, but must accept the outcome.",
	}
}

func (e *RabbleRousing) Name() string {
	return e.name
}

func (e *RabbleRousing) Description() string {
	return e.description
}

func (e *RabbleRousing) Applier() domain.Applier {
	return e.Apply
}

func (e *RabbleRousing) Apply(state domain.GameState) domain.GameState {
	// - Whenever you fail a Guile or Leadership Test, you may re-roll to generate a better result, but must accept the outcome.
	log.Println("applying Rabble-Rousing")
	return state
}

var _ domain.Effect = &RabbleRousing{}
