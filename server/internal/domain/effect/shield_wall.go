package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type ShieldWall struct {
	name        string
	description string
}

func NewShieldWall() *ShieldWall {
	return &ShieldWall{
		name:        "Shield Wall",
		description: "When Engaged with an ally and they fail to Parry or cannot do so, you may immediately Parry in their stead for 1 AP. If successful, they suffer no Damage (and neither do you). You must have a shield in-hand in order to use this Trait.",
	}
}

func (e *ShieldWall) Name() string {
	return e.name
}

func (e *ShieldWall) Description() string {
	return e.description
}

func (e *ShieldWall) Applier() domain.Applier {
	return e.Apply
}

func (e *ShieldWall) Apply(state domain.GameState) domain.GameState {
	// - When Engaged with an ally and they fail to Parry or cannot do so, you may immediately Parry in their stead for 1 AP. If successful, they suffer no Damage (and neither do you). You must have a shield in-hand in order to use this Trait.
	log.Println("applying Shield Wall")
	return state
}

var _ domain.Effect = &ShieldWall{}
