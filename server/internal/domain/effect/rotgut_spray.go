package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type RotgutSpray struct {
  name string
  description string
}

func NewRotgutSpray() *RotgutSpray {
  return &RotgutSpray{
    name: "Rotgut Spray",
    description: "Effect1",
  }
}

func (e *RotgutSpray) Name() string {
  return e.name
}

func (e *RotgutSpray) Description() string {
  return e.description
}

func (e *RotgutSpray) Applier() domain.Applier {
  return e.Apply
}

func (e *RotgutSpray) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Rotgut Spray")
  return state
}

var _ domain.Effect = &RotgutSpray{}
