package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type HeadTraumaUntil struct {
	name        string
	description string
}

func NewHeadTraumaUntil() *HeadTraumaUntil {
	return &HeadTraumaUntil{
		name:        "Head Trauma Until",
		description: "fully Recuperated, you cannot use Special Actions in combat.",
	}
}

func (e *HeadTraumaUntil) Name() string {
	return e.name
}

func (e *HeadTraumaUntil) Description() string {
	return e.description
}

func (e *HeadTraumaUntil) Applier() domain.Applier {
	return e.Apply
}

func (e *HeadTraumaUntil) Apply(state domain.GameState) domain.GameState {
	// - fully Recuperated, you cannot use Special Actions in combat.
	log.Println("applying Head Trauma Until")
	return state
}

var _ domain.Effect = &HeadTraumaUntil{}
