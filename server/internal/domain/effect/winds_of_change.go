package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type WindsOfChange struct {
	name        string
	description string
}

func NewWindsOfChange() *WindsOfChange {
	return &WindsOfChange{
		name:        "Winds Of Change",
		description: "When using Burst, Cone and Explosion Template Magick, you can reshape it so as to not harm your allies, hurting only your foes.",
	}
}

func (e *WindsOfChange) Name() string {
	return e.name
}

func (e *WindsOfChange) Description() string {
	return e.description
}

func (e *WindsOfChange) Applier() domain.Applier {
	return e.Apply
}

func (e *WindsOfChange) Apply(state domain.GameState) domain.GameState {
	// - When using Burst, Cone and Explosion Template Magick, you can reshape it so as to not harm your allies, hurting only your foes.
	log.Println("applying Winds Of Change")
	return state
}

var _ domain.Effect = &WindsOfChange{}
