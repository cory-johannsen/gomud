package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type ItsaTrap struct {
	name        string
	description string
}

func NewItsaTrap() *ItsaTrap {
	return &ItsaTrap{
		name:        "It's a Trap!",
		description: "You automatically spot hidden traps without having to make a Skill Test. Furthermore, you may flip the results to succeed at Skill Tests to create and disarm traps.",
	}
}

func (e *ItsaTrap) Name() string {
	return e.name
}

func (e *ItsaTrap) Description() string {
	return e.description
}

func (e *ItsaTrap) Applier() domain.Applier {
	return e.Apply
}

func (e *ItsaTrap) Apply(state domain.GameState) domain.GameState {
	// - You automatically spot hidden traps without having to make a Skill Test. Furthermore, you may flip the results to succeed at Skill Tests to create and disarm traps.
	log.Println("applying It's a Trap!")
	return state
}

var _ domain.Effect = &ItsaTrap{}
