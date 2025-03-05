package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type SwordBoard struct {
	name        string
	description string
}

func NewSwordBoard() *SwordBoard {
	return &SwordBoard{
		name:        "Sword & Board",
		description: "If a foe successfully Parries your Melee Attack, immediately make an Opportunity Attack using a shield as an improvised hand weapon.",
	}
}

func (e *SwordBoard) Name() string {
	return e.name
}

func (e *SwordBoard) Description() string {
	return e.description
}

func (e *SwordBoard) Applier() domain.Applier {
	return e.Apply
}

func (e *SwordBoard) Apply(state domain.GameState) domain.GameState {
	// - If a foe successfully Parries your Melee Attack, immediately make an Opportunity Attack using a shield as an improvised hand weapon.
	log.Println("applying Sword & Board")
	return state
}

var _ domain.Effect = &SwordBoard{}
