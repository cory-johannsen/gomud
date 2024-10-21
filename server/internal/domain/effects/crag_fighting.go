package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type CragFighting struct {
  Name string
  Description string
}

func (e *CragFighting) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Crag Fighting")
  return state
}
