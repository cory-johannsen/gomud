package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type WardenoftheWild struct {
  Name string
  Description string
}

func (e *WardenoftheWild) Apply(state domain.State) domain.State {
  // - When you fail an Awareness or Handle Animal Test, you may re-roll to generate a better result, but must accept the outcome.
  log.Println("applying Warden of the Wild")
  return state
}
