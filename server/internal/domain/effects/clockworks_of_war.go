package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ClockworksofWar struct {
  Name string
  Description string
}

func (e *ClockworksofWar) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Clockworks of War")
  return state
}
