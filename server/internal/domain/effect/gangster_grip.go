package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type GangsterGrip struct {
	name        string
	description string
}

func NewGangsterGrip() *GangsterGrip {
	return &GangsterGrip{
		name:        "Gangster Grip",
		description: "When you make an Attack Action with a weapon possessing the Gunpowder Quality within one yard of an opponent, you inflict an additional 1D6 Fury Die to Damage.",
	}
}

func (e *GangsterGrip) Name() string {
	return e.name
}

func (e *GangsterGrip) Description() string {
	return e.description
}

func (e *GangsterGrip) Applier() domain.Applier {
	return e.Apply
}

func (e *GangsterGrip) Apply(state domain.GameState) domain.GameState {
	// - When you make an Attack Action with a weapon possessing the Gunpowder Quality within one yard of an opponent, you inflict an additional 1D6 Fury Die to Damage.
	log.Println("applying Gangster Grip")
	return state
}

var _ domain.Effect = &GangsterGrip{}
