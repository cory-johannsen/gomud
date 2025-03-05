package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type HereComesTheCavalry struct {
	name        string
	description string
}

func NewHereComesTheCavalry() *HereComesTheCavalry {
	return &HereComesTheCavalry{
		name:        "Here Comes The Cavalry",
		description: "When you fail a Handle Animal or Ride Test, you may re-roll to generate a better result, but must accept the outcome. In addition, when using the Movement subtype of Ride, you do not have to add the additional 1 AP cost.",
	}
}

func (e *HereComesTheCavalry) Name() string {
	return e.name
}

func (e *HereComesTheCavalry) Description() string {
	return e.description
}

func (e *HereComesTheCavalry) Applier() domain.Applier {
	return e.Apply
}

func (e *HereComesTheCavalry) Apply(state domain.GameState) domain.GameState {
	// - When you fail a Handle Animal or Ride Test, you may re-roll to generate a better result, but must accept the outcome. In addition, when using the Movement subtype of Ride, you do not have to add the additional 1 AP cost.
	log.Println("applying Here Comes The Cavalry")
	return state
}

var _ domain.Effect = &HereComesTheCavalry{}
