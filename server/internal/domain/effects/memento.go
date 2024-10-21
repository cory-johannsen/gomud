package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Memento struct {
  Name string
  Description string
}

func (e *Memento) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Memento")
  return state
}
