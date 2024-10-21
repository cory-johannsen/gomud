package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type SaltyDog struct {
  Name string
  Description string
}

func (e *SaltyDog) Apply(state domain.State) domain.State {
  // - After you make an Attack Action with a melee weapon possessing the Finesse Quality, immediately make an Opportunity Attack with any one-handed ranged weapon on the same Turn. The weapon must already be loaded or in-hand.
  log.Println("applying Salty Dog")
  return state
}
