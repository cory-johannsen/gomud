package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Bewitching struct {
  Name string
  Description string
}

func (e *Bewitching) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Bewitching")
  return state
}
