package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Underfoot struct {
	name        string
	description string
}

func NewUnderfoot() *Underfoot {
	return &Underfoot{
		name:        "Underfoot",
		description: "Add +3 to your Initiative.",
	}
}

func (e *Underfoot) Name() string {
	return e.name
}

func (e *Underfoot) Description() string {
	return e.description
}

func (e *Underfoot) Applier() domain.Applier {
	return e.Apply
}

func (e *Underfoot) Apply(state domain.GameState) domain.GameState {
	// - Add +3 to your Initiative.
	log.Println("applying Underfoot")
	return state
}

var _ domain.Effect = &Underfoot{}
