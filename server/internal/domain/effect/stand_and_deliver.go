package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type StandAndDeliver struct {
	name        string
	description string
}

func NewStandAndDeliver() *StandAndDeliver {
	return &StandAndDeliver{
		name:        "Stand And Deliver",
		description: "At your option, you may substitute the Charm Skill in place of Intimidate. In addition, you may substitute Charm in place of any Skill required to Parry melee weapons.",
	}
}

func (e *StandAndDeliver) Name() string {
	return e.name
}

func (e *StandAndDeliver) Description() string {
	return e.description
}

func (e *StandAndDeliver) Applier() domain.Applier {
	return e.Apply
}

func (e *StandAndDeliver) Apply(state domain.GameState) domain.GameState {
	// - At your option, you may substitute the Charm Skill in place of Intimidate. In addition, you may substitute Charm in place of any Skill required to Parry melee weapons.
	log.Println("applying Stand And Deliver")
	return state
}

var _ domain.Effect = &StandAndDeliver{}
