package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Protective struct {
	name        string
	description string
}

func NewProtective() *Protective {
	return &Protective{
		name:        "Protective",
		description: "Shields of this Quality may be used to Parry any Attack Action made with a ranged weapon.",
	}
}

func (e *Protective) Name() string {
	return e.name
}

func (e *Protective) Description() string {
	return e.description
}

func (e *Protective) Applier() domain.Applier {
	return e.Apply
}

func (e *Protective) Apply(state domain.GameState) domain.GameState {
	// - Shields of this Quality may be used to Parry any Attack Action made with a ranged weapon.
	log.Println("applying Protective")
	return state
}

var _ domain.Effect = &Protective{}
