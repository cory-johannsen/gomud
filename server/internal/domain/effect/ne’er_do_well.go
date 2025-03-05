package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type NeerDoWell struct {
	name        string
	description string
}

func NewNeerDoWell() *NeerDoWell {
	return &NeerDoWell{
		name:        "Ne’er Do Well",
		description: "You cannot Assist others’ Skill Tests.",
	}
}

func (e *NeerDoWell) Name() string {
	return e.name
}

func (e *NeerDoWell) Description() string {
	return e.description
}

func (e *NeerDoWell) Applier() domain.Applier {
	return e.Apply
}

func (e *NeerDoWell) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
	log.Println("applying Ne’er Do Well")
	return state
}

var _ domain.Effect = &NeerDoWell{}
