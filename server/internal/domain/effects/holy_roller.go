package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type HolyRoller struct {
  Name string
  Description string
}

func (e *HolyRoller) Apply(state domain.State) domain.State {
  // - Adjust your Damage Threshold by +3, but only when you aren’t wearing armor.
  log.Println("applying Holy Roller")
  return state
}
