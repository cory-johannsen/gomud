package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type GrimResolve struct {
  Name string
  Description string
}

func (e *GrimResolve) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Grim Resolve")
  return state
}
