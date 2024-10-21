package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Oathkeeper struct {
  Name string
  Description string
}

func (e *Oathkeeper) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Oathkeeper")
  return state
}
