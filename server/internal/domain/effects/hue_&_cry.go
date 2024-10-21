package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type HueCry struct {
  Name string
  Description string
}

func (e *HueCry) Apply(state domain.State) domain.State {
  // - When combat begins, roll 2D10, instead of 1D10, to determine your place in the Initiative Order.
  log.Println("applying Hue & Cry")
  return state
}
