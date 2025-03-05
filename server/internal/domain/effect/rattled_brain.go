package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type RattledBrain struct {
	name        string
	description string
}

func NewRattledBrain() *RattledBrain {
	return &RattledBrain{
		name:        "Rattled Brain",
		description: "Until fully Recuperated, reduce your Initiative by 3.",
	}
}

func (e *RattledBrain) Name() string {
	return e.name
}

func (e *RattledBrain) Description() string {
	return e.description
}

func (e *RattledBrain) Applier() domain.Applier {
	return e.Apply
}

func (e *RattledBrain) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, reduce your Initiative by 3.
	log.Println("applying Rattled Brain")
	return state
}

var _ domain.Effect = &RattledBrain{}
