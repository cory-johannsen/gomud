package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type SacredMantra struct {
	name        string
	description string
}

func NewSacredMantra() *SacredMantra {
	return &SacredMantra{
		name:        "Sacred Mantra",
		description: "You may enter a sacred trance for one hour. If you succeed at a Resolve Test at the end of the trance, you expel all Intoxication and Poisons from your system.",
	}
}

func (e *SacredMantra) Name() string {
	return e.name
}

func (e *SacredMantra) Description() string {
	return e.description
}

func (e *SacredMantra) Applier() domain.Applier {
	return e.Apply
}

func (e *SacredMantra) Apply(state domain.GameState) domain.GameState {
	// - You may enter a sacred trance for one hour. If you succeed at a Resolve Test at the end of the trance, you expel all Intoxication and Poisons from your system.
	log.Println("applying Sacred Mantra")
	return state
}

var _ domain.Effect = &SacredMantra{}
