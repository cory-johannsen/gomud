package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Handspring struct {
	name        string
	description string
}

func NewHandspring() *Handspring {
	return &Handspring{
		name:        "Handspring",
		description: "You can use the Movement Action called Get Up for 0 Action Points.",
	}
}

func (e *Handspring) Name() string {
	return e.name
}

func (e *Handspring) Description() string {
	return e.description
}

func (e *Handspring) Applier() domain.Applier {
	return e.Apply
}

func (e *Handspring) Apply(state domain.GameState) domain.GameState {
	// - You can use the Movement Action called Get Up for 0 Action Points.
	log.Println("applying Handspring")
	return state
}

var _ domain.Effect = &Handspring{}
