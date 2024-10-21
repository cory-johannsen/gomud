package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type NaturesOwn struct {
  Name string
  Description string
}

func (e *NaturesOwn) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Nature's Own")
  return state
}
