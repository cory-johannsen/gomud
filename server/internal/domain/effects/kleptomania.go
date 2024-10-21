package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Kleptomania struct {
  Name string
  Description string
}

func (e *Kleptomania) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Kleptomania")
  return state
}
