package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Light struct {
	name        string
	description string
}

func NewLight() *Light {
	return &Light{
		name:        "Light",
		description: "Whenever weapons or shields of this Quality are held in your off-hand when attacking with a melee weapon in your primary hand, add +1 to Total Damage.",
	}
}

func (e *Light) Name() string {
	return e.name
}

func (e *Light) Description() string {
	return e.description
}

func (e *Light) Applier() domain.Applier {
	return e.Apply
}

func (e *Light) Apply(state domain.GameState) domain.GameState {
	// - Whenever weapons or shields of this Quality are held in your off-hand when attacking with a melee weapon in your primary hand, add +1 to Total Damage.
	log.Println("applying Light")
	return state
}

var _ domain.Effect = &Light{}
