package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Siegecraft struct {
	name        string
	description string
}

func NewSiegecraft() *Siegecraft {
	return &Siegecraft{
		name:        "Siegecraft",
		description: "When you attempt to determine appropriate distances for siege engines and employ them to hit, you gain a +20 Base Chance to Warfare Tests.",
	}
}

func (e *Siegecraft) Name() string {
	return e.name
}

func (e *Siegecraft) Description() string {
	return e.description
}

func (e *Siegecraft) Applier() domain.Applier {
	return e.Apply
}

func (e *Siegecraft) Apply(state domain.GameState) domain.GameState {
	// - When you attempt to determine appropriate distances for siege engines and employ them to hit, you gain a +20 Base Chance to Warfare Tests.
	log.Println("applying Siegecraft")
	return state
}

var _ domain.Effect = &Siegecraft{}
