package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Chatty struct {
  Name string
  Description string
}

func (e *Chatty) Apply(state domain.State) domain.State {
  // - "When using Flair to solve a problem, flip the roll to Success.  Otherwise flip the roll to Failure."
  log.Println("applying Chatty")
  return state
}
