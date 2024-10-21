package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type AimlessDrifter struct {
  Name string
  Description string
}

func (e *AimlessDrifter) Apply(state domain.State) domain.State {
  // - When you fail a Folklore or Navigation Test, you may re-roll to generate a better result, but must accept the outcome.
  log.Println("applying Aimless Drifter")
  return state
}
