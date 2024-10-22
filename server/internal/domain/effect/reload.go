package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Reload struct {
  name string
  description string
}

func NewReload() *Reload {
  return &Reload{
    name: "Reload",
    description: "The weapon can be reloaded after it is empty.",
  }
}

func (e *Reload) Name() string {
  return e.name
}

func (e *Reload) Description() string {
  return e.description
}

func (e *Reload) Applier() domain.Applier {
  return e.Apply
}

func (e *Reload) Apply(state domain.State) domain.State {
  // - The weapon can be reloaded after it is empty.
  log.Println("applying Reload")
  return state
}

var _ domain.Effect = &Reload{}
