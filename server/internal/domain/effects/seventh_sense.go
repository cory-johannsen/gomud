package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type SeventhSense struct {
  Name string
  Description string
}

func (e *SeventhSense) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Seventh Sense")
  return state
}
