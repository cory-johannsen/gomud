package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MortalityWork struct {
  Name string
  Description string
}

func (e *MortalityWork) Apply(state domain.State) domain.State {
  // - When you inflict an Injury with a weapon, your foe also suffers 2D10+your [WB] mental Peril.
  log.Println("applying Mortality Work")
  return state
}
