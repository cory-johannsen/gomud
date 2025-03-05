package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Dunderhead struct {
	name        string
	description string
}

func NewDunderhead() *Dunderhead {
	return &Dunderhead{
		name:        "Dunderhead",
		description: "Whenever you suffer mental Peril, move one additional step down the Peril Condition Track negatively while suffering 1 Corruption.",
	}
}

func (e *Dunderhead) Name() string {
	return e.name
}

func (e *Dunderhead) Description() string {
	return e.description
}

func (e *Dunderhead) Applier() domain.Applier {
	return e.Apply
}

func (e *Dunderhead) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
	log.Println("applying Dunderhead")
	return state
}

var _ domain.Effect = &Dunderhead{}
