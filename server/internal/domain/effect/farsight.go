package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Farsight struct {
  name string
  description string
}

func NewFarsight() *Farsight {
  return &Farsight{
    name: "Farsight",
    description: "When attempting to discern minute details with your vision, you always succeed at Awareness Tests.,",
  }
}

func (e *Farsight) Name() string {
  return e.name
}

func (e *Farsight) Description() string {
  return e.description
}

func (e *Farsight) Applier() domain.Applier {
  return e.Apply
}

func (e *Farsight) Apply(state domain.State) domain.State {
  // - When attempting to discern minute details with your vision, you always succeed at Awareness Tests.,
  log.Println("applying Farsight")
  return state
}

var _ domain.Effect = &Farsight{}
