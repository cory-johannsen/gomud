package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type CholericTemperament struct {
	name        string
	description string
}

func NewCholericTemperament() *CholericTemperament {
	return &CholericTemperament{
		name:        "Choleric Temperament",
		description: "Whenever you roll Chaos Dice to determine if you Injure a foe and fail to do so, move one step down the Peril Condition Track negatively while suffering 1 Corruption.",
	}
}

func (e *CholericTemperament) Name() string {
	return e.name
}

func (e *CholericTemperament) Description() string {
	return e.description
}

func (e *CholericTemperament) Applier() domain.Applier {
	return e.Apply
}

func (e *CholericTemperament) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
	log.Println("applying Choleric Temperament")
	return state
}

var _ domain.Effect = &CholericTemperament{}
