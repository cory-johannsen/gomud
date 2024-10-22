package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type StrengthoftheMountain struct {
  name string
  description string
}

func NewStrengthoftheMountain() *StrengthoftheMountain {
  return &StrengthoftheMountain{
    name: "Strength of the Mountain",
    description: "Effect1",
  }
}

func (e *StrengthoftheMountain) Name() string {
  return e.name
}

func (e *StrengthoftheMountain) Description() string {
  return e.description
}

func (e *StrengthoftheMountain) Applier() domain.Applier {
  return e.Apply
}

func (e *StrengthoftheMountain) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Strength of the Mountain")
  return state
}

var _ domain.Effect = &StrengthoftheMountain{}