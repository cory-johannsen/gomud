package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type DogsofWar struct {
  Name string
  Description string
}

func (e *DogsofWar) Apply(state domain.State) domain.State {
  // - You never suffer Serious Injuries.
  log.Println("applying Dogs of War")
  return state
}
