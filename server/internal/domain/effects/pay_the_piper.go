package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type PaythePiper struct {
  Name string
  Description string
}

func (e *PaythePiper) Apply(state domain.State) domain.State {
  // - Whenever you fail a Bargain or Intimidate Test, you may re-roll to generate a better result, but must accept the outcome.
  log.Println("applying Pay the Piper")
  return state
}
