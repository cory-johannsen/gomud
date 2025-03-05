package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type MinorConcussion struct {
	name        string
	description string
}

func NewMinorConcussion() *MinorConcussion {
	return &MinorConcussion{
		name:        "Minor Concussion",
		description: "Until fully Recuperated, you remain Incapacitated!",
	}
}

func (e *MinorConcussion) Name() string {
	return e.name
}

func (e *MinorConcussion) Description() string {
	return e.description
}

func (e *MinorConcussion) Applier() domain.Applier {
	return e.Apply
}

func (e *MinorConcussion) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you remain Incapacitated!
	log.Println("applying Minor Concussion")
	return state
}

var _ domain.Effect = &MinorConcussion{}
