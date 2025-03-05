package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type TownGossip struct {
	name        string
	description string
}

func NewTownGossip() *TownGossip {
	return &TownGossip{
		name:        "Town Gossip",
		description: "When you fail an Eavesdrop or Rumor Test, you may re-roll to generate a better result, but must accept the outcome.",
	}
}

func (e *TownGossip) Name() string {
	return e.name
}

func (e *TownGossip) Description() string {
	return e.description
}

func (e *TownGossip) Applier() domain.Applier {
	return e.Apply
}

func (e *TownGossip) Apply(state domain.GameState) domain.GameState {
	// - When you fail an Eavesdrop or Rumor Test, you may re-roll to generate a better result, but must accept the outcome.
	log.Println("applying Town Gossip")
	return state
}

var _ domain.Effect = &TownGossip{}
