package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Grudgebearer struct {
  Name string
  Description string
}

func (e *Grudgebearer) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Grudgebearer")
  return state
}
