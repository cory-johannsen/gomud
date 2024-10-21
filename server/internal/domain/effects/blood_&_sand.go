package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type BloodSand struct {
  Name string
  Description string
}

func (e *BloodSand) Apply(state domain.State) domain.State {
  // - Whenever you spend a Fortune Point, you move one step up the Damage & Peril Condition Tracks positively.
  log.Println("applying Blood & Sand")
  return state
}
