package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MeditativeHealing struct {
  Name string
  Description string
}

func (e *MeditativeHealing) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Meditative Healing")
  return state
}
