package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Immolate struct {
  name string
  description string
}

func NewImmolate() *Immolate {
  return &Immolate{
    name: "Immolate",
    description: "After a foe is struck with a weapon of this Quality, they must succeed at a Coordination Test or be set On Fire. The weapon is immediately extinguished after a successful attack, until relit.",
  }
}

func (e *Immolate) Name() string {
  return e.name
}

func (e *Immolate) Description() string {
  return e.description
}

func (e *Immolate) Applier() domain.Applier {
  return e.Apply
}

func (e *Immolate) Apply(state domain.State) domain.State {
  // - After a foe is struck with a weapon of this Quality, they must succeed at a Coordination Test or be set On Fire. The weapon is immediately extinguished after a successful attack, until relit.
  log.Println("applying Immolate")
  return state
}

var _ domain.Effect = &Immolate{}
