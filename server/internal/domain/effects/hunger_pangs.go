package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type HungerPangs struct {
  Name string
  Description string
}

func (e *HungerPangs) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Hunger Pangs")
  return state
}
