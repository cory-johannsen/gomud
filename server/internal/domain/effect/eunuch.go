package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Eunuch struct {
	name        string
	description string
}

func NewEunuch() *Eunuch {
	return &Eunuch{
		name:        "Eunuch",
		description: "You are immune to the charms and seduction by those who find you attractive, and unable to have children. However, being made victim to these same charms and seduction causes you to suffer 1 Corruption.",
	}
}

func (e *Eunuch) Name() string {
	return e.name
}

func (e *Eunuch) Description() string {
	return e.description
}

func (e *Eunuch) Applier() domain.Applier {
	return e.Apply
}

func (e *Eunuch) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
	log.Println("applying Eunuch")
	return state
}

var _ domain.Effect = &Eunuch{}
