package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Pintsized struct {
  name string
  description string
}

func NewPintsized() *Pintsized {
  return &Pintsized{
    name: "Pintsized",
    description: "Effect1",
  }
}

func (e *Pintsized) Name() string {
  return e.name
}

func (e *Pintsized) Description() string {
  return e.description
}

func (e *Pintsized) Applier() domain.Applier {
  return e.Apply
}

func (e *Pintsized) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Pintsized")
  return state
}

var _ domain.Effect = &Pintsized{}
