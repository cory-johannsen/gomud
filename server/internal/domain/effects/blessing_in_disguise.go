package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type BlessinginDisguise struct {
  Name string
  Description string
}

func (e *BlessinginDisguise) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Blessing in Disguise")
  return state
}
