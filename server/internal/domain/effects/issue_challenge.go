package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type IssueChallenge struct {
  Name string
  Description string
}

func (e *IssueChallenge) Apply(state domain.State) domain.State {
  // - When Engaged with an ally and they fail to Parry or cannot do so, you may immediately Parry in their stead for 1 AP. If successful, they suffer no Damage (and neither do you). You must have a shield in-hand in order to use this Trait.
  log.Println("applying Issue Challenge")
  return state
}
