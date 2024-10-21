package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Murdery struct {
  Name string
  Description string
}

func (e *Murdery) Apply(state domain.State) domain.State {
  // - "When performing an act of Killing to solve a problem, flip the roll to Success.  Otherwise flip the roll to Failure."
  log.Println("applying Murdery")
  return state
}
