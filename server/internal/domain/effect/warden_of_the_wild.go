package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type WardenoftheWild struct {
	name        string
	description string
}

func NewWardenoftheWild() *WardenoftheWild {
	return &WardenoftheWild{
		name:        "Warden of the Wild",
		description: "When you fail an Awareness or Handle Animal Test, you may re-roll to generate a better result, but must accept the outcome.",
	}
}

func (e *WardenoftheWild) Name() string {
	return e.name
}

func (e *WardenoftheWild) Description() string {
	return e.description
}

func (e *WardenoftheWild) Applier() domain.Applier {
	return e.Apply
}

func (e *WardenoftheWild) Apply(state domain.GameState) domain.GameState {
	// - When you fail an Awareness or Handle Animal Test, you may re-roll to generate a better result, but must accept the outcome.
	log.Println("applying Warden of the Wild")
	return state
}

var _ domain.Effect = &WardenoftheWild{}
