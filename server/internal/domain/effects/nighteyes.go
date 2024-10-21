package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Nighteyes struct {
  Name string
  Description string
}

func (e *Nighteyes) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Nighteyes")
  return state
}
