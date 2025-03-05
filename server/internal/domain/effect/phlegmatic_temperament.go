package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type PhlegmaticTemperament struct {
	name        string
	description string
}

func NewPhlegmaticTemperament() *PhlegmaticTemperament {
	return &PhlegmaticTemperament{
		name:        "Phlegmatic Temperament",
		description: "Whenever you are suffering from Stress, Fear or Terror, your Fury Dice do not explode. This lasts until you get a good night's rest.",
	}
}

func (e *PhlegmaticTemperament) Name() string {
	return e.name
}

func (e *PhlegmaticTemperament) Description() string {
	return e.description
}

func (e *PhlegmaticTemperament) Applier() domain.Applier {
	return e.Apply
}

func (e *PhlegmaticTemperament) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
	log.Println("applying Phlegmatic Temperament")
	return state
}

var _ domain.Effect = &PhlegmaticTemperament{}
