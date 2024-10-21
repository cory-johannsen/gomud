package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type StressFractureUntil struct {
  name string
  description string
}

func NewStressFractureUntil() *StressFractureUntil {
  return &StressFractureUntil{
    name: "Stress Fracture Until",
    description: "fully Recuperated, you cannot Counterspell, Dodge or Parry.",
  }
}

func (e *StressFractureUntil) Name() string {
  return e.name
}

func (e *StressFractureUntil) Description() string {
  return e.description
}

func (e *StressFractureUntil) Applier() domain.Applier {
  return e.Apply
}

func (e *StressFractureUntil) Apply(state domain.State) domain.State {
  // - fully Recuperated, you cannot Counterspell, Dodge or Parry.
  log.Println("applying Stress Fracture Until")
  return state
}

var _ domain.Effect = &StressFractureUntil{}
