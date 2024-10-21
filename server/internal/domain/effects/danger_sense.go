package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type DangerSense struct {
  Name string
  Description string
}

func (e *DangerSense) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Danger Sense")
  return state
}
