package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Ironclad struct {
  Name string
  Description string
}

func (e *Ironclad) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Ironclad")
  return state
}
