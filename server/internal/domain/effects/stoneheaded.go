package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Stoneheaded struct {
  Name string
  Description string
}

func (e *Stoneheaded) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Stoneheaded")
  return state
}
