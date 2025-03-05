package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type GrimResolve struct {
	name        string
	description string
}

func NewGrimResolve() *GrimResolve {
	return &GrimResolve{
		name:        "Grim Resolve",
		description: "Immediately after you suffer Damage from a melee or ranged weapon, spend a Fortune Point to ignore it entirely, therefore negating both Damage and any Injuries you may have suffered. You may even use this after failing to Dodge or Parry.,",
	}
}

func (e *GrimResolve) Name() string {
	return e.name
}

func (e *GrimResolve) Description() string {
	return e.description
}

func (e *GrimResolve) Applier() domain.Applier {
	return e.Apply
}

func (e *GrimResolve) Apply(state domain.GameState) domain.GameState {
	// - Immediately after you suffer Damage from a melee or ranged weapon, spend a Fortune Point to ignore it entirely, therefore negating both Damage and any Injuries you may have suffered. You may even use this after failing to Dodge or Parry.,
	log.Println("applying Grim Resolve")
	return state
}

var _ domain.Effect = &GrimResolve{}
