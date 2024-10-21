package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type FameFortune struct {
  Name string
  Description string
}

func (e *FameFortune) Apply(state domain.State) domain.State {
  // - When making any Skill Test, you never suffer the ill-effects of Critical Failures, instead treating it as a failed Skill Test.
  log.Println("applying Fame & Fortune")
  return state
}
