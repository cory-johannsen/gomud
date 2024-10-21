package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Indifference struct {
  name string
  description string
}

func NewIndifference() *Indifference {
  return &Indifference{
    name: "Indifference",
    description: "When viewing blood-soaked and visceral scenes of death, you always succeed at Resolve Tests to withstand their potential psychological effect (such as with Stress, Fear and Terror) and don’t suffer Corruption due to such scenes.",
  }
}

func (e *Indifference) Name() string {
  return e.name
}

func (e *Indifference) Description() string {
  return e.description
}

func (e *Indifference) Applier() domain.Applier {
  return e.Apply
}

func (e *Indifference) Apply(state domain.State) domain.State {
  // - When viewing blood-soaked and visceral scenes of death, you always succeed at Resolve Tests to withstand their potential psychological effect (such as with Stress, Fear and Terror) and don’t suffer Corruption due to such scenes.
  log.Println("applying Indifference")
  return state
}

var _ domain.Effect = &Indifference{}
