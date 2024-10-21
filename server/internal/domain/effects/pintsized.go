package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Pintsized struct {
  Name string
  Description string
}

func (e *Pintsized) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Pintsized")
  return state
}
