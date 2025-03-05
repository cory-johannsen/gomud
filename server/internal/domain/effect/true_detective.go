package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type TrueDetective struct {
	name        string
	description string
}

func NewTrueDetective() *TrueDetective {
	return &TrueDetective{
		name:        "True Detective",
		description: "When Intoxicated or underneath the effect of Deliriants, make a Scrutinize Test. If successful, you can ask for the GM to give you an important clue from your investigations you may not have already thought of or overlooked. This benefit cannot be gained more than once per day.",
	}
}

func (e *TrueDetective) Name() string {
	return e.name
}

func (e *TrueDetective) Description() string {
	return e.description
}

func (e *TrueDetective) Applier() domain.Applier {
	return e.Apply
}

func (e *TrueDetective) Apply(state domain.GameState) domain.GameState {
	// - When Intoxicated or underneath the effect of Deliriants, make a Scrutinize Test. If successful, you can ask for the GM to give you an important clue from your investigations you may not have already thought of or overlooked. This benefit cannot be gained more than once per day.
	log.Println("applying True Detective")
	return state
}

var _ domain.Effect = &TrueDetective{}
