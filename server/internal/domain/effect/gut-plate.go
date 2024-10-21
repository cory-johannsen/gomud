package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Gutplate struct {
  name string
  description string
}

func NewGutplate() *Gutplate {
  return &Gutplate{
    name: "Gut-plate",
    description: "Effect1",
  }
}

func (e *Gutplate) Name() string {
  return e.name
}

func (e *Gutplate) Description() string {
  return e.description
}

func (e *Gutplate) Applier() domain.Applier {
  return e.Apply
}

func (e *Gutplate) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Gut-plate")
  return state
}

var _ domain.Effect = &Gutplate{}
