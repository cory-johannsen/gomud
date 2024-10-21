package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MightyThews struct {
  Name string
  Description string
}

func (e *MightyThews) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Mighty Thews")
  return state
}
