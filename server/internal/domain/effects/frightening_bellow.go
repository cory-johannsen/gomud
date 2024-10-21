package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type FrighteningBellow struct {
  Name string
  Description string
}

func (e *FrighteningBellow) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Frightening Bellow")
  return state
}
