package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type FrighteningBellow struct {
  name string
  description string
}

func NewFrighteningBellow() *FrighteningBellow {
  return &FrighteningBellow{
    name: "Frightening Bellow",
    description: "Effect1",
  }
}

func (e *FrighteningBellow) Name() string {
  return e.name
}

func (e *FrighteningBellow) Description() string {
  return e.description
}

func (e *FrighteningBellow) Applier() domain.Applier {
  return e.Apply
}

func (e *FrighteningBellow) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Frightening Bellow")
  return state
}

var _ domain.Effect = &FrighteningBellow{}
