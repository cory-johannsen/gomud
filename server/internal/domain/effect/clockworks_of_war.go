package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ClockworksofWar struct {
  name string
  description string
}

func NewClockworksofWar() *ClockworksofWar {
  return &ClockworksofWar{
    name: "Clockworks of War",
    description: "Provided you have the appropriate tools, any two- handed weapon can be modified into a compact one-handed version (which only you can wield). Despite being one-handed, it retains its original function, but now has an Encumbrance of 1.,",
  }
}

func (e *ClockworksofWar) Name() string {
  return e.name
}

func (e *ClockworksofWar) Description() string {
  return e.description
}

func (e *ClockworksofWar) Applier() domain.Applier {
  return e.Apply
}

func (e *ClockworksofWar) Apply(state domain.State) domain.State {
  // - Provided you have the appropriate tools, any two- handed weapon can be modified into a compact one-handed version (which only you can wield). Despite being one-handed, it retains its original function, but now has an Encumbrance of 1.,
  log.Println("applying Clockworks of War")
  return state
}

var _ domain.Effect = &ClockworksofWar{}
