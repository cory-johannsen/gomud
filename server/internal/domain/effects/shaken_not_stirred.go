package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ShakenNotStirred struct {
  Name string
  Description string
}

func (e *ShakenNotStirred) Apply(state domain.State) domain.State {
  // - When you fail a Charm or Eavesdrop Test, you may re-roll to generate a better result, but must accept the outcome.
  log.Println("applying Shaken Not Stirred")
  return state
}
