package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type FourinHand struct {
  Name string
  Description string
}

func (e *FourinHand) Apply(state domain.State) domain.State {
  // - You may flip the results to succeed at Drive Tests. When you succeed, it is always considered a Critical Success. In addition, when using the Movement subtype of Drive, you do not have to add the additional 1 AP cost.
  log.Println("applying Four-in-Hand")
  return state
}
