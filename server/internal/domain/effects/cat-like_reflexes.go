package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type CatlikeReflexes struct {
  Name string
  Description string
}

func (e *CatlikeReflexes) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Cat-like Reflexes")
  return state
}
