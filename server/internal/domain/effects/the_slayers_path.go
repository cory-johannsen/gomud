package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type TheSlayersPath struct {
  Name string
  Description string
}

func (e *TheSlayersPath) Apply(state domain.State) domain.State {
  // - After you have successfully defeated a creature (providing it is not a Humanoid or of a player Ancestry), its kind then becomes your chosen enemy. When attacking your chosen enemy, you always add a 1D6 Fury Die to melee weapon attacks. In addition, you are immune to the Intimidate Skill – along with Stress, Fear and Terror – used by these creatures. Note that there is no upper limit to the number of chosen enemy types a Slayer can have.
  log.Println("applying The Slayers Path")
  return state
}
