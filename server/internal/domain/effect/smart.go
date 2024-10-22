package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Smart struct {
	name        string
	description string
}

func NewSmart() *Smart {
	return &Smart{
		name:        "Smart",
		description: "'When using Reasoning to solve a problem, flip the roll to Success.  Otherwise flip the roll to Failure.'",
	}
}

func (e *Smart) Name() string {
	return e.name
}

func (e *Smart) Description() string {
	return e.description
}

func (e *Smart) Applier() domain.Applier {
	return e.Apply
}

func (e *Smart) Apply(state domain.State) domain.State {
	// - 'When using Reasoning to solve a problem, flip the roll to Success.  Otherwise flip the roll to Failure.'
	log.Println("applying Smart")
	return state
}

var _ domain.Effect = &Smart{}
