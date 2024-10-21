package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ItsaTrap struct {
  Name string
  Description string
}

func (e *ItsaTrap) Apply(state domain.State) domain.State {
  // - You automatically spot hidden traps without having to make a Skill Test. Furthermore, you may flip the results to succeed at Skill Tests to create and disarm traps.
  log.Println("applying It's a Trap!")
  return state
}
