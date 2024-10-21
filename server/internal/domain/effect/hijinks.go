package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Hijinks struct {
  name string
  description string
}

func NewHijinks() *Hijinks {
  return &Hijinks{
    name: "Hijinks",
    description: "Effect1",
  }
}

func (e *Hijinks) Name() string {
  return e.name
}

func (e *Hijinks) Description() string {
  return e.description
}

func (e *Hijinks) Applier() domain.Applier {
  return e.Apply
}

func (e *Hijinks) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Hijinks")
  return state
}

var _ domain.Effect = &Hijinks{}
