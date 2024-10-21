package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type EscapeArtist struct {
  Name string
  Description string
}

func (e *EscapeArtist) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Escape Artist")
  return state
}
