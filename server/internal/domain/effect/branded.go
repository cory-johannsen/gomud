package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Branded struct {
	name        string
	description string
}

func NewBranded() *Branded {
	return &Branded{
		name:        "Branded",
		description: "Whenever you interact with those who know you’re Branded and hold you in contempt due to it, you cannot succeed at Fellowship-based Skill Tests to interact with them.",
	}
}

func (e *Branded) Name() string {
	return e.name
}

func (e *Branded) Description() string {
	return e.description
}

func (e *Branded) Applier() domain.Applier {
	return e.Apply
}

func (e *Branded) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
	log.Println("applying Branded")
	return state
}

var _ domain.Effect = &Branded{}
