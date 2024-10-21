package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ChildrenoftheEarth struct {
  Name string
  Description string
}

func (e *ChildrenoftheEarth) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Children of the Earth")
  return state
}
