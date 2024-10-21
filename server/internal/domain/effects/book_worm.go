package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type BookWorm struct {
  Name string
  Description string
}

func (e *BookWorm) Apply(state domain.State) domain.State {
  // - Multiply your [IB] by three to determine how many Focuses you may possess.
  log.Println("applying Book Worm")
  return state
}
