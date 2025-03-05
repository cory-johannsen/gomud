package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Knifework struct {
	name        string
	description string
}

func NewKnifework() *Knifework {
	return &Knifework{
		name:        "Knifework",
		description: "When you deal Damage to a foe using a melee weapon with the Fast Quality, they must Resist with a Toughness Test or begin to Bleed.",
	}
}

func (e *Knifework) Name() string {
	return e.name
}

func (e *Knifework) Description() string {
	return e.description
}

func (e *Knifework) Applier() domain.Applier {
	return e.Apply
}

func (e *Knifework) Apply(state domain.GameState) domain.GameState {
	// - When you deal Damage to a foe using a melee weapon with the Fast Quality, they must Resist with a Toughness Test or begin to Bleed.
	log.Println("applying Knifework")
	return state
}

var _ domain.Effect = &Knifework{}
