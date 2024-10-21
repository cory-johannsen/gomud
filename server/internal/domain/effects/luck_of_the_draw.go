package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type LuckoftheDraw struct {
  Name string
  Description string
}

func (e *LuckoftheDraw) Apply(state domain.State) domain.State {
  // - When you use Fortune Points, you do not need to roll percentile dice to make a Skill Test. Instead, you automatically Critically Succeed at the Skill Test you intended to make.
  log.Println("applying Luck of the Draw")
  return state
}
