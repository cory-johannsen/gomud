package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type NobleSavage struct {
  Name string
  Description string
}

func (e *NobleSavage) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Noble Savage")
  return state
}
