package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type TornShoulder struct {
  name string
  description string
}

func NewTornShoulder() *TornShoulder {
  return &TornShoulder{
    name: "Torn Shoulder",
    description: "Whatever you are holding in your primary hand gains the Ruined! Quality. Until fully Recuperated, you start your Turn with 2 less AP.",
  }
}

func (e *TornShoulder) Name() string {
  return e.name
}

func (e *TornShoulder) Description() string {
  return e.description
}

func (e *TornShoulder) Applier() domain.Applier {
  return e.Apply
}

func (e *TornShoulder) Apply(state domain.State) domain.State {
  // - Whatever you are holding in your primary hand gains the Ruined! Quality. Until fully Recuperated, you start your Turn with 2 less AP.
  log.Println("applying Torn Shoulder")
  return state
}

var _ domain.Effect = &TornShoulder{}
