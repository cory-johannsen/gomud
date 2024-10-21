package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Recidivist struct {
  Name string
  Description string
}

func (e *Recidivist) Apply(state domain.State) domain.State {
  // - "When performing an Illegal Act to solve a problem, flip the roll to Success.  Otherwise flip the roll to Failure."
  log.Println("applying Recidivist")
  return state
}
