package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type LowBlow struct {
  Name string
  Description string
}

func (e *LowBlow) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Low Blow")
  return state
}
