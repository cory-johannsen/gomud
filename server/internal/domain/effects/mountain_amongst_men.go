package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MountainAmongstMen struct {
  Name string
  Description string
}

func (e *MountainAmongstMen) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Mountain Amongst Men")
  return state
}
