package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type TheShowMustGoOn struct {
  Name string
  Description string
}

func (e *TheShowMustGoOn) Apply(state domain.State) domain.State {
  // - You can attempt a Skill Test, related to your Focus in the arts, to perform for others. If successful, a number of allies equal to your [FB] gain the benefits of Inspiring Words for 24 hours. If you Critically Succeed, you influence a number of allies equal to three times your [FB] instead. A Character may only gain benefit of this performance once per day, which takes at least an hour of time to perform.
  log.Println("applying The Show Must Go On")
  return state
}
