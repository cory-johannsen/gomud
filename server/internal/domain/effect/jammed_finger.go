package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type JammedFinger struct {
	name        string
	description string
}

func NewJammedFinger() *JammedFinger {
	return &JammedFinger{
		name:        "Jammed Finger",
		description: "You immediately drop whatever you are holding. Until fully Recuperated, you must flip the results to fail all melee weapon attacks with your primary hand.",
	}
}

func (e *JammedFinger) Name() string {
	return e.name
}

func (e *JammedFinger) Description() string {
	return e.description
}

func (e *JammedFinger) Applier() domain.Applier {
	return e.Apply
}

func (e *JammedFinger) Apply(state domain.GameState) domain.GameState {
	// - You immediately drop whatever you are holding. Until fully Recuperated, you must flip the results to fail all melee weapon attacks with your primary hand.
	log.Println("applying Jammed Finger")
	return state
}

var _ domain.Effect = &JammedFinger{}
