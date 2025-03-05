package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Entangling struct {
	name        string
	description string
}

func NewEntangling() *Entangling {
	return &Entangling{
		name:        "Entangling",
		description: "Immediately after striking a foe, weapons of this Quality force a foe to Resist either a Chokehold or Takedown – you choose which. Additionally, whenever a foe is threatened with a Chokehold or Takedown with this weapon, they must flip the results to fail when Resisting its effects.",
	}
}

func (e *Entangling) Name() string {
	return e.name
}

func (e *Entangling) Description() string {
	return e.description
}

func (e *Entangling) Applier() domain.Applier {
	return e.Apply
}

func (e *Entangling) Apply(state domain.GameState) domain.GameState {
	// - Immediately after striking a foe, weapons of this Quality force a foe to Resist either a Chokehold or Takedown – you choose which. Additionally, whenever a foe is threatened with a Chokehold or Takedown with this weapon, they must flip the results to fail when Resisting its effects.
	log.Println("applying Entangling")
	return state
}

var _ domain.Effect = &Entangling{}
