package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type DeadlyAim struct {
	name        string
	description string
}

func NewDeadlyAim() *DeadlyAim {
	return &DeadlyAim{
		name:        "Deadly Aim",
		description: "Whenever you make an Attack Action or Perilous Stunt using a ranged weapon, foes must flip the results to fail to Dodge, Parry or Resist it.,",
	}
}

func (e *DeadlyAim) Name() string {
	return e.name
}

func (e *DeadlyAim) Description() string {
	return e.description
}

func (e *DeadlyAim) Applier() domain.Applier {
	return e.Apply
}

func (e *DeadlyAim) Apply(state domain.GameState) domain.GameState {
	// - Whenever you make an Attack Action or Perilous Stunt using a ranged weapon, foes must flip the results to fail to Dodge, Parry or Resist it.,
	log.Println("applying Deadly Aim")
	return state
}

var _ domain.Effect = &DeadlyAim{}
