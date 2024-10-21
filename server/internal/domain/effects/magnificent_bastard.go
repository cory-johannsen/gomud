package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MagnificentBastard struct {
  Name string
  Description string
}

func (e *MagnificentBastard) Apply(state domain.State) domain.State {
  // - Select one foe in combat. That foe is left Defenseless to all Perilous Stunts you make until they are defeated. You may select another foe afterwards.
  log.Println("applying Magnificent Bastard")
  return state
}
