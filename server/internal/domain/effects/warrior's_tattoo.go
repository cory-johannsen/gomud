package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type WarriorsTattoo struct {
  Name string
  Description string
}

func (e *WarriorsTattoo) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Warrior's Tattoo")
  return state
}
