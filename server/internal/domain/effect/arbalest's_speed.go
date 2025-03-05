package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type ArbalestsSpeed struct {
	name        string
	description string
}

func NewArbalestsSpeed() *ArbalestsSpeed {
	return &ArbalestsSpeed{
		name:        "Arbalest's Speed",
		description: "When you begin to Load a ranged weapon, make a successful Coordination Test to Load it without spending Action Points. If you fail, you must spend the requisite number of Action Points to finish Loading.",
	}
}

func (e *ArbalestsSpeed) Name() string {
	return e.name
}

func (e *ArbalestsSpeed) Description() string {
	return e.description
}

func (e *ArbalestsSpeed) Applier() domain.Applier {
	return e.Apply
}

func (e *ArbalestsSpeed) Apply(state domain.GameState) domain.GameState {
	// - When you begin to Load a ranged weapon, make a successful Coordination Test to Load it without spending Action Points. If you fail, you must spend the requisite number of Action Points to finish Loading.
	log.Println("applying Arbalest's Speed")
	return state
}

var _ domain.Effect = &ArbalestsSpeed{}
