package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type BeyondtheVeil struct {
  Name string
  Description string
}

func (e *BeyondtheVeil) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Beyond the Veil")
  return state
}
