package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type HereComesTheCavalry struct {
  Name string
  Description string
}

func (e *HereComesTheCavalry) Apply(state domain.State) domain.State {
  // - When you fail a Handle Animal or Ride Test, you may re-roll to generate a better result, but must accept the outcome. In addition, when using the Movement subtype of Ride, you do not have to add the additional 1 AP cost.
  log.Println("applying Here Comes The Cavalry")
  return state
}
