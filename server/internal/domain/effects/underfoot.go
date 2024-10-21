package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Underfoot struct {
  Name string
  Description string
}

func (e *Underfoot) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Underfoot")
  return state
}
