package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Mariner struct {
  name string
  description string
}

func NewMariner() *Mariner {
  return &Mariner{
    name: "Mariner",
    description: "When boating in waters where the shore can be seen, you gain a +20 Base Chance to Pilot Tests.",
  }
}

func (e *Mariner) Name() string {
  return e.name
}

func (e *Mariner) Description() string {
  return e.description
}

func (e *Mariner) Applier() domain.Applier {
  return e.Apply
}

func (e *Mariner) Apply(state domain.State) domain.State {
  // - When boating in waters where the shore can be seen, you gain a +20 Base Chance to Pilot Tests.
  log.Println("applying Mariner")
  return state
}

var _ domain.Effect = &Mariner{}
