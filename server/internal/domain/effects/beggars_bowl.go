package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type BeggarsBowl struct {
  Name string
  Description string
}

func (e *BeggarsBowl) Apply(state domain.State) domain.State {
  // - Whenever you fail a Guile or Intimidate Test, you may re-roll to generate a better result, but must accept the outcome.
  log.Println("applying Beggars Bowl")
  return state
}
