package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Reach struct {
  name string
  description string
}

func NewReach() *Reach {
  return &Reach{
    name: "Reach",
    description: "Weapons of this Quality may strike a foe you’re Engaged with or standing one yard away from outside of an Engagement. Additionally, foes who are armed with a Reach weapon can make an Opportunity Attack whenever someone Charges or Runs toward them.",
  }
}

func (e *Reach) Name() string {
  return e.name
}

func (e *Reach) Description() string {
  return e.description
}

func (e *Reach) Applier() domain.Applier {
  return e.Apply
}

func (e *Reach) Apply(state domain.State) domain.State {
  // - Weapons of this Quality may strike a foe you’re Engaged with or standing one yard away from outside of an Engagement. Additionally, foes who are armed with a Reach weapon can make an Opportunity Attack whenever someone Charges or Runs toward them.
  log.Println("applying Reach")
  return state
}

var _ domain.Effect = &Reach{}
