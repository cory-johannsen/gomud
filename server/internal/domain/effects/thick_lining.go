package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ThickLining struct {
  Name string
  Description string
}

func (e *ThickLining) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Thick Lining")
  return state
}
