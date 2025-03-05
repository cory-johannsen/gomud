package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Gatecrasher struct {
	name        string
	description string
}

func NewGatecrasher() *Gatecrasher {
	return &Gatecrasher{
		name:        "Gatecrasher",
		description: "When you Take Aim and then make a successful Melee Attack, add 3 Damage.",
	}
}

func (e *Gatecrasher) Name() string {
	return e.name
}

func (e *Gatecrasher) Description() string {
	return e.description
}

func (e *Gatecrasher) Applier() domain.Applier {
	return e.Apply
}

func (e *Gatecrasher) Apply(state domain.GameState) domain.GameState {
	// - When you Take Aim and then make a successful Melee Attack, add 3 Damage.
	log.Println("applying Gatecrasher")
	return state
}

var _ domain.Effect = &Gatecrasher{}
