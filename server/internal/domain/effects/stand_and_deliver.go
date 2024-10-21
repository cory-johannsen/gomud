package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type StandAndDeliver struct {
  Name string
  Description string
}

func (e *StandAndDeliver) Apply(state domain.State) domain.State {
  // - At your option, you may substitute the Charm Skill in place of Intimidate. In addition, you may substitute Charm in place of any Skill required to Parry melee weapons.
  log.Println("applying Stand And Deliver")
  return state
}
