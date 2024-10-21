package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type DeadlyAim struct {
  Name string
  Description string
}

func (e *DeadlyAim) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Deadly Aim")
  return state
}
