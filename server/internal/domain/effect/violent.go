package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Violent struct {
	name        string
	description string
}

func NewViolent() *Violent {
	return &Violent{
		name:        "Violent",
		description: "Effect1",
	}
}

func (e *Violent) Name() string {
	return e.name
}

func (e *Violent) Description() string {
	return e.description
}

func (e *Violent) Applier() domain.Applier {
	return e.Apply
}

func (e *Violent) Apply(state domain.GameState) domain.GameState {
	// - Effect1
	log.Println("applying Violent")
	return state
}

var _ domain.Effect = &Violent{}
