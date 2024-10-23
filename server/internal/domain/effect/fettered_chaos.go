package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type FetteredChaos struct {
  name string
  description string
}

func NewFetteredChaos() *FetteredChaos {
  return &FetteredChaos{
    name: "Fettered Chaos",
    description: "Whenever you suffer Corruption, decrease the number you gain by three (to a minimum of one). This means that if you suffer 6 Corruption, you gain 3 instead.,",
  }
}

func (e *FetteredChaos) Name() string {
  return e.name
}

func (e *FetteredChaos) Description() string {
  return e.description
}

func (e *FetteredChaos) Applier() domain.Applier {
  return e.Apply
}

func (e *FetteredChaos) Apply(state domain.State) domain.State {
  // - Whenever you suffer Corruption, decrease the number you gain by three (to a minimum of one). This means that if you suffer 6 Corruption, you gain 3 instead.,
  log.Println("applying Fettered Chaos")
  return state
}

var _ domain.Effect = &FetteredChaos{}
