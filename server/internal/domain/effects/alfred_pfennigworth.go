package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type AlfredPfennigworth struct {
  Name string
  Description string
}

func (e *AlfredPfennigworth) Apply(state domain.State) domain.State {
  // - When you Assist others with Skill Tests, they gain two Assist Dice instead, choosing the better of the two to use. In addition, when you Assist them, they may ignore any Critical Failures they roll, treating them as regular failures instead.
  log.Println("applying Alfred Pfennigworth")
  return state
}
