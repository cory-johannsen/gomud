package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type PulledMuscle struct {
	name        string
	description string
}

func NewPulledMuscle() *PulledMuscle {
	return &PulledMuscle{
		name:        "Pulled Muscle",
		description: "Until fully Recuperated, you must flip the results to fail all Brawn-based Skill Tests.",
	}
}

func (e *PulledMuscle) Name() string {
	return e.name
}

func (e *PulledMuscle) Description() string {
	return e.description
}

func (e *PulledMuscle) Applier() domain.Applier {
	return e.Apply
}

func (e *PulledMuscle) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you must flip the results to fail all Brawn-based Skill Tests.
	log.Println("applying Pulled Muscle")
	return state
}

var _ domain.Effect = &PulledMuscle{}
