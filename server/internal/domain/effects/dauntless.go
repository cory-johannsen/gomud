package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Dauntless struct {
  Name string
  Description string
}

func (e *Dauntless) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Dauntless")
  return state
}
