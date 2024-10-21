package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MortalCombat struct {
  Name string
  Description string
}

func (e *MortalCombat) Apply(state domain.State) domain.State {
  // - You ignore the Pummeling and Weak Qualities when fighting with Brawling type of weapons. In addition, you may refer to [BB] or [CB] when inflicting Damage with this type of weapon, whichever is more favorable.
  log.Println("applying Mortal Combat!")
  return state
}
