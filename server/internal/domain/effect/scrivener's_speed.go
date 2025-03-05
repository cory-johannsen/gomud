package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type ScrivenersSpeed struct {
	name        string
	description string
}

func NewScrivenersSpeed() *ScrivenersSpeed {
	return &ScrivenersSpeed{
		name:        "Scrivener's Speed",
		description: "You reduce the time required to use Education Tests by half (examples include research, speed-reading, writing, etc.). In addition, you automatically succeed at all Skill Tests to decipher cryptographs in languages they understand.",
	}
}

func (e *ScrivenersSpeed) Name() string {
	return e.name
}

func (e *ScrivenersSpeed) Description() string {
	return e.description
}

func (e *ScrivenersSpeed) Applier() domain.Applier {
	return e.Apply
}

func (e *ScrivenersSpeed) Apply(state domain.GameState) domain.GameState {
	// - You reduce the time required to use Education Tests by half (examples include research, speed-reading, writing, etc.). In addition, you automatically succeed at all Skill Tests to decipher cryptographs in languages they understand.
	log.Println("applying Scrivener's Speed")
	return state
}

var _ domain.Effect = &ScrivenersSpeed{}
