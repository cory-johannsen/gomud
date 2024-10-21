package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Craven struct {
  Name string
  Description string
}

func (e *Craven) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Craven")
  return state
}
