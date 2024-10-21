package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type RuneMarkedGlory struct {
  Name string
  Description string
}

func (e *RuneMarkedGlory) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Rune-Marked Glory")
  return state
}
