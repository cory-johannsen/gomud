package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type PhysicalProwess struct {
  Name string
  Description string
}

func (e *PhysicalProwess) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Physical Prowess")
  return state
}
