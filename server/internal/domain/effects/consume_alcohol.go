package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ConsumeAlcohol struct {
  Name string
  Description string
}

func (e *ConsumeAlcohol) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Consume Alcohol")
  return state
}
