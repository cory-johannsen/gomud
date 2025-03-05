package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type GruesomeShot struct {
	name        string
	description string
}

func NewGruesomeShot() *GruesomeShot {
	return &GruesomeShot{
		name:        "Gruesome Shot",
		description: "When you Take Aim and then make a successful Ranged Attack, add 3 Damage.",
	}
}

func (e *GruesomeShot) Name() string {
	return e.name
}

func (e *GruesomeShot) Description() string {
	return e.description
}

func (e *GruesomeShot) Applier() domain.Applier {
	return e.Apply
}

func (e *GruesomeShot) Apply(state domain.GameState) domain.GameState {
	// - When you Take Aim and then make a successful Ranged Attack, add 3 Damage.
	log.Println("applying Gruesome Shot")
	return state
}

var _ domain.Effect = &GruesomeShot{}
