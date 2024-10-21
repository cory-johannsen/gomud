package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type SacredMantra struct {
  Name string
  Description string
}

func (e *SacredMantra) Apply(state domain.State) domain.State {
  // - You may enter a sacred trance for one hour. If you succeed at a Resolve Test at the end of the trance, you expel all Intoxication and Poisons from your system.
  log.Println("applying Sacred Mantra")
  return state
}
