package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type TheShowMustGoOn struct {
  name string
  description string
}

func NewTheShowMustGoOn() *TheShowMustGoOn {
  return &TheShowMustGoOn{
    name: "The Show Must Go On",
    description: "You can attempt a Skill Test, related to your Focus in the arts, to perform for others. If successful, a number of allies equal to your [FB] gain the benefits of Inspiring Words for 24 hours. If you Critically Succeed, you influence a number of allies equal to three times your [FB] instead. A Character may only gain benefit of this performance once per day, which takes at least an hour of time to perform.",
  }
}

func (e *TheShowMustGoOn) Name() string {
  return e.name
}

func (e *TheShowMustGoOn) Description() string {
  return e.description
}

func (e *TheShowMustGoOn) Applier() domain.Applier {
  return e.Apply
}

func (e *TheShowMustGoOn) Apply(state domain.State) domain.State {
  // - You can attempt a Skill Test, related to your Focus in the arts, to perform for others. If successful, a number of allies equal to your [FB] gain the benefits of Inspiring Words for 24 hours. If you Critically Succeed, you influence a number of allies equal to three times your [FB] instead. A Character may only gain benefit of this performance once per day, which takes at least an hour of time to perform.
  log.Println("applying The Show Must Go On")
  return state
}

var _ domain.Effect = &TheShowMustGoOn{}
