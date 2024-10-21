package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type DungeonsDeep struct {
  Name string
  Description string
}

func (e *DungeonsDeep) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Dungeons Deep")
  return state
}
