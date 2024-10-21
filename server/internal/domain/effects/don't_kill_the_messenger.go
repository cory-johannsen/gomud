package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type DontKilltheMessenger struct {
  Name string
  Description string
}

func (e *DontKilltheMessenger) Apply(state domain.State) domain.State {
  // - When using Fellowship-based Skill Tests, you do not suffer any additional penalties due to differences in Social Class.
  log.Println("applying Don't Kill the Messenger")
  return state
}
