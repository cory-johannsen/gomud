package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Heavy struct {
	name        string
	description string
}

func NewHeavy() *Heavy {
	return &Heavy{
		name:        "Heavy",
		description: "Armor of this Quality prohibits the use of the Incantation Skill to cast Magick and Coordination in order to Dodge attacks.",
	}
}

func (e *Heavy) Name() string {
	return e.name
}

func (e *Heavy) Description() string {
	return e.description
}

func (e *Heavy) Applier() domain.Applier {
	return e.Apply
}

func (e *Heavy) Apply(state domain.GameState) domain.GameState {
	// - Armor of this Quality prohibits the use of the Incantation Skill to cast Magick and Coordination in order to Dodge attacks.
	log.Println("applying Heavy")
	return state
}

var _ domain.Effect = &Heavy{}
