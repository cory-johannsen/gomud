package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type FortunesWheel struct {
	name        string
	description string
}

func NewFortunesWheel() *FortunesWheel {
	return &FortunesWheel{
		name:        "Fortune's Wheel",
		description: "Anytime you generate a Critical Failure after rolling D100, automatically add one Fortune Point into the Fortune Pool.,",
	}
}

func (e *FortunesWheel) Name() string {
	return e.name
}

func (e *FortunesWheel) Description() string {
	return e.description
}

func (e *FortunesWheel) Applier() domain.Applier {
	return e.Apply
}

func (e *FortunesWheel) Apply(state domain.GameState) domain.GameState {
	// - Anytime you generate a Critical Failure after rolling D100, automatically add one Fortune Point into the Fortune Pool.,
	log.Println("applying Fortune's Wheel")
	return state
}

var _ domain.Effect = &FortunesWheel{}
