package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Coordination struct {
  Name string
  Description string
}

func (e *Coordination) Apply(state domain.State) domain.State {
  // -
  log.Println("applying Coordination")
  return state
}
