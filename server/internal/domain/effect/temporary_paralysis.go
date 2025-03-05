package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type TemporaryParalysis struct {
	name        string
	description string
}

func NewTemporaryParalysis() *TemporaryParalysis {
	return &TemporaryParalysis{
		name:        "Temporary Paralysis",
		description: "You are knocked Prone. Until fully Recuperated, you cannot move as you're paralyzed.",
	}
}

func (e *TemporaryParalysis) Name() string {
	return e.name
}

func (e *TemporaryParalysis) Description() string {
	return e.description
}

func (e *TemporaryParalysis) Applier() domain.Applier {
	return e.Apply
}

func (e *TemporaryParalysis) Apply(state domain.GameState) domain.GameState {
	// - You are knocked Prone. Until fully Recuperated, you cannot move as you're paralyzed.
	log.Println("applying Temporary Paralysis")
	return state
}

var _ domain.Effect = &TemporaryParalysis{}
