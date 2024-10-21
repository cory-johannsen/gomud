package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type RoadtoEnlightenment struct {
  Name string
  Description string
}

func (e *RoadtoEnlightenment) Apply(state domain.State) domain.State {
  // - When you suffer 1 to 3 Corruption, make a (Routine +10%) Resolve Test. If you suffer 4 to 6 Corruption, make a (Standard +/-0%) Resolve Test. If you suffer 7 to 9 Corruption, make a (Challenging -10%) Resolve Test. If successful, you suffer no Corruption.
  log.Println("applying Road to Enlightenment")
  return state
}
