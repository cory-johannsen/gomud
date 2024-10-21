package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type DirtySecret struct {
  Name string
  Description string
}

func (e *DirtySecret) Apply(state domain.State) domain.State {
  // - "Whenever the topic that you keep secret comes up, make a Grit roll.  On a Success, you keep your cool.  On a Failure, you Panic."
  log.Println("applying Dirty Secret")
  return state
}
