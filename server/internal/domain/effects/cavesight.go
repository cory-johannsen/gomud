package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Cavesight struct {
  Name string
  Description string
}

func (e *Cavesight) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Cavesight")
  return state
}
