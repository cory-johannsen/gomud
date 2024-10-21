package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Violent struct {
  Name string
  Description string
}

func (e *Violent) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Violent")
  return state
}
