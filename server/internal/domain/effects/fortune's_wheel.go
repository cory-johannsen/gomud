package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type FortunesWheel struct {
  Name string
  Description string
}

func (e *FortunesWheel) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Fortune's Wheel")
  return state
}
