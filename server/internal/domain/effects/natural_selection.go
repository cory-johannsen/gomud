package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type NaturalSelection struct {
  Name string
  Description string
}

func (e *NaturalSelection) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Natural Selection")
  return state
}
