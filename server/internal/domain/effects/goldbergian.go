package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Goldbergian struct {
  Name string
  Description string
}

func (e *Goldbergian) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Goldbergian")
  return state
}
