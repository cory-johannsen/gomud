package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Beguiler struct {
  Name string
  Description string
}

func (e *Beguiler) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Beguiler")
  return state
}
