package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MettleSteel struct {
  Name string
  Description string
}

func (e *MettleSteel) Apply(state domain.State) domain.State {
  // - At your option, you may substitute the Warfare Skill in place of Resolve. In addition, you may substitute Warfare in place of any Skill required to Resist Perilous Stunts.
  log.Println("applying Mettle & Steel")
  return state
}
