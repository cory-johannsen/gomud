package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Farsight struct {
  Name string
  Description string
}

func (e *Farsight) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Farsight")
  return state
}
