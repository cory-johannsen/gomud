package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ConsumeAlcohol struct {
  name string
  description string
}

func NewConsumeAlcohol() *ConsumeAlcohol {
  return &ConsumeAlcohol{
    name: "Consume Alcohol",
    description: "Effect1",
  }
}

func (e *ConsumeAlcohol) Name() string {
  return e.name
}

func (e *ConsumeAlcohol) Description() string {
  return e.description
}

func (e *ConsumeAlcohol) Applier() domain.Applier {
  return e.Apply
}

func (e *ConsumeAlcohol) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Consume Alcohol")
  return state
}

var _ domain.Effect = &ConsumeAlcohol{}
