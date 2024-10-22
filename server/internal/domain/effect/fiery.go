package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Fiery struct {
  name string
  description string
}

func NewFiery() *Fiery {
  return &Fiery{
    name: "Fiery",
    description: "After a foe is struck with a weapon of this Quality, both the foe and one other combatant who is Engaged with them must succeed at a Coordination Test or be set On Fire (see Chapter 9",
  }
}

func (e *Fiery) Name() string {
  return e.name
}

func (e *Fiery) Description() string {
  return e.description
}

func (e *Fiery) Applier() domain.Applier {
  return e.Apply
}

func (e *Fiery) Apply(state domain.State) domain.State {
  // - After a foe is struck with a weapon of this Quality, both the foe and one other combatant who is Engaged with them must succeed at a Coordination Test or be set On Fire (see Chapter 9
  log.Println("applying Fiery")
  return state
}

var _ domain.Effect = &Fiery{}
