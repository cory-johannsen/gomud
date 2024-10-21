package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type WarriorsTattoo struct {
  name string
  description string
}

func NewWarriorsTattoo() *WarriorsTattoo {
  return &WarriorsTattoo{
    name: "Warrior's Tattoo",
    description: "Effect1",
  }
}

func (e *WarriorsTattoo) Name() string {
  return e.name
}

func (e *WarriorsTattoo) Description() string {
  return e.description
}

func (e *WarriorsTattoo) Applier() domain.Applier {
  return e.Apply
}

func (e *WarriorsTattoo) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Warrior's Tattoo")
  return state
}

var _ domain.Effect = &WarriorsTattoo{}
