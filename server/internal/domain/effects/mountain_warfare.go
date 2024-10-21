package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MountainWarfare struct {
  Name string
  Description string
}

func (e *MountainWarfare) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Mountain Warfare")
  return state
}
