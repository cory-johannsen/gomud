package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type TortuousInquisition struct {
  Name string
  Description string
}

func (e *TortuousInquisition) Apply(state domain.State) domain.State {
  // - You do not suffer the ill-effects of Peril, until you are at “Ignore 3 Skill Ranks” on the Peril Condition Track.
  log.Println("applying Tortuous Inquisition")
  return state
}
