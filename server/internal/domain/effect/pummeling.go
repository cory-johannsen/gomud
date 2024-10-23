package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Pummeling struct {
  name string
  description string
}

func NewPummeling() *Pummeling {
  return &Pummeling{
    name: "Pummeling",
    description: "Weapons of this Quality always refer to [MB] whenever inflicting Damage, instead of [BB]. Pummeling weapons can only inflict Moderate Injuries, never Serious or Grievous Injuries. Finally, Pummeling weapons cannot be used to cause Bleeding.",
  }
}

func (e *Pummeling) Name() string {
  return e.name
}

func (e *Pummeling) Description() string {
  return e.description
}

func (e *Pummeling) Applier() domain.Applier {
  return e.Apply
}

func (e *Pummeling) Apply(state domain.State) domain.State {
  // - Weapons of this Quality always refer to [MB] whenever inflicting Damage, instead of [BB]. Pummeling weapons can only inflict Moderate Injuries, never Serious or Grievous Injuries. Finally, Pummeling weapons cannot be used to cause Bleeding.
  log.Println("applying Pummeling")
  return state
}

var _ domain.Effect = &Pummeling{}
