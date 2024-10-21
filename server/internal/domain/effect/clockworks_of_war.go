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
    description: "Effect1",
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
  // - Effect1
  log.Println("applying Clockworks of War")
  return state
}

var _ domain.Effect = &ClockworksofWar{}
