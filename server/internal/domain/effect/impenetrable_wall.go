package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type ImpenetrableWall struct {
	name        string
	description string
}

func NewImpenetrableWall() *ImpenetrableWall {
	return &ImpenetrableWall{
		name:        "Impenetrable Wall",
		description: "Opponents do not gain an advantage (such as additional Damage or bonuses to strike) when they flank, outnumber or surround you in combat.",
	}
}

func (e *ImpenetrableWall) Name() string {
	return e.name
}

func (e *ImpenetrableWall) Description() string {
	return e.description
}

func (e *ImpenetrableWall) Applier() domain.Applier {
	return e.Apply
}

func (e *ImpenetrableWall) Apply(state domain.GameState) domain.GameState {
	// - Opponents do not gain an advantage (such as additional Damage or bonuses to strike) when they flank, outnumber or surround you in combat.
	log.Println("applying Impenetrable Wall")
	return state
}

var _ domain.Effect = &ImpenetrableWall{}
