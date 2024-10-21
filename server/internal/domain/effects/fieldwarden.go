package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Fieldwarden struct {
  Name string
  Description string
}

func (e *Fieldwarden) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Fieldwarden")
  return state
}
