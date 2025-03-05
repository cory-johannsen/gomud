package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type SplitFace struct {
	name        string
	description string
}

func NewSplitFace() *SplitFace {
	return &SplitFace{
		name:        "Split Face",
		description: "You must flip the results to fail all Skill Tests which rely on smell and taste.",
	}
}

func (e *SplitFace) Name() string {
	return e.name
}

func (e *SplitFace) Description() string {
	return e.description
}

func (e *SplitFace) Applier() domain.Applier {
	return e.Apply
}

func (e *SplitFace) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
	log.Println("applying Split Face")
	return state
}

var _ domain.Effect = &SplitFace{}
