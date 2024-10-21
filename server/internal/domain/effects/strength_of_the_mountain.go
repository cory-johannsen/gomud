package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type StrengthoftheMountain struct {
  Name string
  Description string
}

func (e *StrengthoftheMountain) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Strength of the Mountain")
  return state
}
