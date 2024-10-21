package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Wendigo struct {
  Name string
  Description string
}

func (e *Wendigo) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Wendigo")
  return state
}
