package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type HocusPocus struct {
  Name string
  Description string
}

func (e *HocusPocus) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Hocus Pocus")
  return state
}
