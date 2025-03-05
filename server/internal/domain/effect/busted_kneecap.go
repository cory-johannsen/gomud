package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type BustedKneecap struct {
	name        string
	description string
}

func NewBustedKneecap() *BustedKneecap {
	return &BustedKneecap{
		name:        "Busted Kneecap",
		description: "Until fully Recuperated, any time you fail a Skill Test that relies on Brawn or Agility, you suffer 2D10+2 physical Peril.",
	}
}

func (e *BustedKneecap) Name() string {
	return e.name
}

func (e *BustedKneecap) Description() string {
	return e.description
}

func (e *BustedKneecap) Applier() domain.Applier {
	return e.Apply
}

func (e *BustedKneecap) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, any time you fail a Skill Test that relies on Brawn or Agility, you suffer 2D10+2 physical Peril.
	log.Println("applying Busted Kneecap")
	return state
}

var _ domain.Effect = &BustedKneecap{}
