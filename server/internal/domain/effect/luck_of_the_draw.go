package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type LuckoftheDraw struct {
	name        string
	description string
}

func NewLuckoftheDraw() *LuckoftheDraw {
	return &LuckoftheDraw{
		name:        "Luck of the Draw",
		description: "When you use Fortune Points, you do not need to roll percentile dice to make a Skill Test. Instead, you automatically Critically Succeed at the Skill Test you intended to make.",
	}
}

func (e *LuckoftheDraw) Name() string {
	return e.name
}

func (e *LuckoftheDraw) Description() string {
	return e.description
}

func (e *LuckoftheDraw) Applier() domain.Applier {
	return e.Apply
}

func (e *LuckoftheDraw) Apply(state domain.GameState) domain.GameState {
	// - When you use Fortune Points, you do not need to roll percentile dice to make a Skill Test. Instead, you automatically Critically Succeed at the Skill Test you intended to make.
	log.Println("applying Luck of the Draw")
	return state
}

var _ domain.Effect = &LuckoftheDraw{}
