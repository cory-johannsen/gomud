package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Nemesis struct {
	name        string
	description string
}

func NewNemesis() *Nemesis {
	return &Nemesis{
		name:        "Nemesis",
		description: "When confronted by your Nemesis, you cannot sacrifice Fate or Fortune Points, as their presence confounds you. Your Nemesis is determined by the GM, who will likely take your ideas about your Character’s history or past rivals into consideration.",
	}
}

func (e *Nemesis) Name() string {
	return e.name
}

func (e *Nemesis) Description() string {
	return e.description
}

func (e *Nemesis) Applier() domain.Applier {
	return e.Apply
}

func (e *Nemesis) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
	log.Println("applying Nemesis")
	return state
}

var _ domain.Effect = &Nemesis{}
