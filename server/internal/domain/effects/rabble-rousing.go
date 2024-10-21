package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type RabbleRousing struct {
  Name string
  Description string
}

func (e *RabbleRousing) Apply(state domain.State) domain.State {
  // - Whenever you fail a Guile or Leadership Test, you may re-roll to generate a better result, but must accept the outcome.
  log.Println("applying Rabble-Rousing")
  return state
}
