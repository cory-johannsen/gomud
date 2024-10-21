package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Firstborn struct {
  Name string
  Description string
}

func (e *Firstborn) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Firstborn")
  return state
}
