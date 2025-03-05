package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type ForAFewShillingsMore struct {
	name        string
	description string
}

func NewForAFewShillingsMore() *ForAFewShillingsMore {
	return &ForAFewShillingsMore{
		name:        "For A Few Shillings More",
		description: "In combat, whenever you make your first Attack Action with a ranged weapon, you never miss. In addition, your intended target cannot Dodge, Parry or Resist the attack. You also add an additional 1D6 Fury Die to the same attack.",
	}
}

func (e *ForAFewShillingsMore) Name() string {
	return e.name
}

func (e *ForAFewShillingsMore) Description() string {
	return e.description
}

func (e *ForAFewShillingsMore) Applier() domain.Applier {
	return e.Apply
}

func (e *ForAFewShillingsMore) Apply(state domain.GameState) domain.GameState {
	// - In combat, whenever you make your first Attack Action with a ranged weapon, you never miss. In addition, your intended target cannot Dodge, Parry or Resist the attack. You also add an additional 1D6 Fury Die to the same attack.
	log.Println("applying For A Few Shillings More")
	return state
}

var _ domain.Effect = &ForAFewShillingsMore{}
