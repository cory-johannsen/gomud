package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type HyperextendedElbow struct {
	name        string
	description string
}

func NewHyperextendedElbow() *HyperextendedElbow {
	return &HyperextendedElbow{
		name:        "Hyperextended Elbow",
		description: "Until fully Recuperated, you must flip the results to fail all Actions in Combat that rely on ranged weapons.",
	}
}

func (e *HyperextendedElbow) Name() string {
	return e.name
}

func (e *HyperextendedElbow) Description() string {
	return e.description
}

func (e *HyperextendedElbow) Applier() domain.Applier {
	return e.Apply
}

func (e *HyperextendedElbow) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you must flip the results to fail all Actions in Combat that rely on ranged weapons.
	log.Println("applying Hyperextended Elbow")
	return state
}

var _ domain.Effect = &HyperextendedElbow{}
