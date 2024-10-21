package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type SituationalAwareness struct {
  Name string
  Description string
}

func (e *SituationalAwareness) Apply(state domain.State) domain.State {
  // - When you fail an Awareness or Stealth Test, you may re-roll to generate a better result, but must accept the outcome.
  log.Println("applying Situational Awareness")
  return state
}
