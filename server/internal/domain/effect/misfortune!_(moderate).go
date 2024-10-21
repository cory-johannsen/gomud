package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MisfortuneModerate struct {
  name string
  description string
}

func NewMisfortuneModerate() *MisfortuneModerate {
  return &MisfortuneModerate{
    name: "Misfortune! (Moderate)",
    description: "Roll on the Serious Injury table instead!",
  }
}

func (e *MisfortuneModerate) Name() string {
  return e.name
}

func (e *MisfortuneModerate) Description() string {
  return e.description
}

func (e *MisfortuneModerate) Applier() domain.Applier {
  return e.Apply
}

func (e *MisfortuneModerate) Apply(state domain.State) domain.State {
  // - Roll on the Serious Injury table instead!
  log.Println("applying Misfortune! (Moderate)")
  return state
}

var _ domain.Effect = &MisfortuneModerate{}
