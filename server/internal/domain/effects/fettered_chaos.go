package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type FetteredChaos struct {
  Name string
  Description string
}

func (e *FetteredChaos) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Fettered Chaos")
  return state
}
