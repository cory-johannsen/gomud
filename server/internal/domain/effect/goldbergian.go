package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Goldbergian struct {
  name string
  description string
}

func NewGoldbergian() *Goldbergian {
  return &Goldbergian{
    name: "Goldbergian",
    description: "Effect1",
  }
}

func (e *Goldbergian) Name() string {
  return e.name
}

func (e *Goldbergian) Description() string {
  return e.description
}

func (e *Goldbergian) Applier() domain.Applier {
  return e.Apply
}

func (e *Goldbergian) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Goldbergian")
  return state
}

var _ domain.Effect = &Goldbergian{}
