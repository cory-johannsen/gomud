package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Repeating struct {
  name string
  description string
}

func NewRepeating() *Repeating {
  return &Repeating{
    name: "Repeating",
    description: "Ranged weapons of this Quality can be fired multiple times without having to spend Action Points to load. See Ammunition.",
  }
}

func (e *Repeating) Name() string {
  return e.name
}

func (e *Repeating) Description() string {
  return e.description
}

func (e *Repeating) Applier() domain.Applier {
  return e.Apply
}

func (e *Repeating) Apply(state domain.State) domain.State {
  // - Ranged weapons of this Quality can be fired multiple times without having to spend Action Points to load. See Ammunition.
  log.Println("applying Repeating")
  return state
}

var _ domain.Effect = &Repeating{}
