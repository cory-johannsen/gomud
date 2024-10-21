package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ArtfulDodger struct {
  Name string
  Description string
}

func (e *ArtfulDodger) Apply(state domain.State) domain.State {
  // - You automatically gain every Focus in the Stealth Skill when you enter this Job. This means you may exceed the normal limits for Focuses set by your [IB], but for Stealth only.
  log.Println("applying Artful Dodger")
  return state
}
