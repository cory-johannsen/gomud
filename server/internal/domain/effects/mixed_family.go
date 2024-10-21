package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MixedFamily struct {
  Name string
  Description string
}

func (e *MixedFamily) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Mixed Family")
  return state
}
