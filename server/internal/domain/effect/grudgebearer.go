package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Grudgebearer struct {
	name        string
	description string
}

func NewGrudgebearer() *Grudgebearer {
	return &Grudgebearer{
		name:        "Grudgebearer",
		description: "Whenever you use Fury Dice to determine weapon Damage, they explode on face ‘5-6’.,",
	}
}

func (e *Grudgebearer) Name() string {
	return e.name
}

func (e *Grudgebearer) Description() string {
	return e.description
}

func (e *Grudgebearer) Applier() domain.Applier {
	return e.Apply
}

func (e *Grudgebearer) Apply(state domain.GameState) domain.GameState {
	// - Whenever you use Fury Dice to determine weapon Damage, they explode on face ‘5-6’.,
	log.Println("applying Grudgebearer")
	return state
}

var _ domain.Effect = &Grudgebearer{}
