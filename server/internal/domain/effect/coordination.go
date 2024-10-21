package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Coordination struct {
  name string
  description string
}

func NewCoordination() *Coordination {
  return &Coordination{
    name: "Coordination",
    description: "",
  }
}

func (e *Coordination) Name() string {
  return e.name
}

func (e *Coordination) Description() string {
  return e.description
}

func (e *Coordination) Applier() domain.Applier {
  return e.Apply
}

func (e *Coordination) Apply(state domain.State) domain.State {
  // -
  log.Println("applying Coordination")
  return state
}

var _ domain.Effect = &Coordination{}
