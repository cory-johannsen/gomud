package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type WretchedPrankster struct {
  Name string
  Description string
}

func (e *WretchedPrankster) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Wretched Prankster")
  return state
}
