package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type OverwhelmingForce struct {
	name        string
	description string
}

func NewOverwhelmingForce() *OverwhelmingForce {
	return &OverwhelmingForce{
		name:        "Overwhelming Force",
		description: "If you roll a Critical Success with a Melee Attack, you inflict the Ruined! Quality to a foe’s shield, weapon or armor (your choice). Overwhelming Force has no effect on Castle-forged trappings, and you must be armed with a two- handed melee weapon.",
	}
}

func (e *OverwhelmingForce) Name() string {
	return e.name
}

func (e *OverwhelmingForce) Description() string {
	return e.description
}

func (e *OverwhelmingForce) Applier() domain.Applier {
	return e.Apply
}

func (e *OverwhelmingForce) Apply(state domain.GameState) domain.GameState {
	// - If you roll a Critical Success with a Melee Attack, you inflict the Ruined! Quality to a foe’s shield, weapon or armor (your choice). Overwhelming Force has no effect on Castle-forged trappings, and you must be armed with a two- handed melee weapon.
	log.Println("applying Overwhelming Force")
	return state
}

var _ domain.Effect = &OverwhelmingForce{}
