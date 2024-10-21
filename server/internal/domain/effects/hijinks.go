package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Hijinks struct {
  Name string
  Description string
}

func (e *Hijinks) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Hijinks")
  return state
}
