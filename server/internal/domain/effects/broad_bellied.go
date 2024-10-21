package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type BroadBellied struct {
  Name string
  Description string
}

func (e *BroadBellied) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Broad Bellied")
  return state
}
