package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type CastIronStomach struct {
  Name string
  Description string
}

func (e *CastIronStomach) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Cast Iron Stomach")
  return state
}
