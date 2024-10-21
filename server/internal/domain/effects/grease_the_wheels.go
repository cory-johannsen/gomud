package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type GreasetheWheels struct {
  Name string
  Description string
}

func (e *GreasetheWheels) Apply(state domain.State) domain.State {
  // - Whenever you fail a Bargain or Counterfeit Test, you may re-roll to generate a better result, but must accept the outcome.
  log.Println("applying Grease the Wheels")
  return state
}
