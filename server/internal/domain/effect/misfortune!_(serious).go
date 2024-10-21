package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MisfortuneSerious struct {
  name string
  description string
}

func NewMisfortuneSerious() *MisfortuneSerious {
  return &MisfortuneSerious{
    name: "Misfortune! (Serious)",
    description: "Roll on the Grievous Injury table instead!",
  }
}

func (e *MisfortuneSerious) Name() string {
  return e.name
}

func (e *MisfortuneSerious) Description() string {
  return e.description
}

func (e *MisfortuneSerious) Applier() domain.Applier {
  return e.Apply
}

func (e *MisfortuneSerious) Apply(state domain.State) domain.State {
  // - Roll on the Grievous Injury table instead!
  log.Println("applying Misfortune! (Serious)")
  return state
}

var _ domain.Effect = &MisfortuneSerious{}
