package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type BeyondtheVeil struct {
  name string
  description string
}

func NewBeyondtheVeil() *BeyondtheVeil {
  return &BeyondtheVeil{
    name: "Beyond the Veil",
    description: "Whenever you fail or Critically Fail an Incantation Test, spend one Fortune Point to turn it into a Critical Success.,",
  }
}

func (e *BeyondtheVeil) Name() string {
  return e.name
}

func (e *BeyondtheVeil) Description() string {
  return e.description
}

func (e *BeyondtheVeil) Applier() domain.Applier {
  return e.Apply
}

func (e *BeyondtheVeil) Apply(state domain.State) domain.State {
  // - Whenever you fail or Critically Fail an Incantation Test, spend one Fortune Point to turn it into a Critical Success.,
  log.Println("applying Beyond the Veil")
  return state
}

var _ domain.Effect = &BeyondtheVeil{}
