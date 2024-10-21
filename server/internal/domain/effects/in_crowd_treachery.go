package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type InCrowdTreachery struct {
  Name string
  Description string
}

func (e *InCrowdTreachery) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying In Crowd Treachery")
  return state
}
