package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type SprainedWrist struct {
	name        string
	description string
}

func NewSprainedWrist() *SprainedWrist {
	return &SprainedWrist{
		name:        "Sprained Wrist",
		description: "Until fully Recuperated, you cannot hold anything in your primary hand, and must rely on your off-hand.",
	}
}

func (e *SprainedWrist) Name() string {
	return e.name
}

func (e *SprainedWrist) Description() string {
	return e.description
}

func (e *SprainedWrist) Applier() domain.Applier {
	return e.Apply
}

func (e *SprainedWrist) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you cannot hold anything in your primary hand, and must rely on your off-hand.
	log.Println("applying Sprained Wrist")
	return state
}

var _ domain.Effect = &SprainedWrist{}
