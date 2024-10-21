package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Mastermind struct {
  Name string
  Description string
}

func (e *Mastermind) Apply(state domain.State) domain.State {
  // - When you fail a Folklore or Intimidate Test, you may re-roll to generate a better result, but must accept the outcome.
  log.Println("applying Mastermind")
  return state
}
