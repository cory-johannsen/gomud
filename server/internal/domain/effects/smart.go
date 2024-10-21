package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Smart struct {
  Name string
  Description string
}

func (e *Smart) Apply(state domain.State) domain.State {
  // - "When using Smarts to solve a problem, flip the roll to Success.  Otherwise flip the roll to Failure."
  log.Println("applying Smart")
  return state
}
