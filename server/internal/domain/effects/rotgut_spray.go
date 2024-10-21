package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type RotgutSpray struct {
  Name string
  Description string
}

func (e *RotgutSpray) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Rotgut Spray")
  return state
}
