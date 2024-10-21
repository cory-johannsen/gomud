package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type DenizenofStone struct {
  Name string
  Description string
}

func (e *DenizenofStone) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Denizen of Stone")
  return state
}
