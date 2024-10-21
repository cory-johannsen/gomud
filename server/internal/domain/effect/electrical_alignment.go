package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ElectricalAlignment struct {
  name string
  description string
}

func NewElectricalAlignment() *ElectricalAlignment {
  return &ElectricalAlignment{
    name: "Electrical Alignment",
    description: "When attempting to Counterhack, you gain a +20 Base Chance to succeed",
  }
}

func (e *ElectricalAlignment) Name() string {
  return e.name
}

func (e *ElectricalAlignment) Description() string {
  return e.description
}

func (e *ElectricalAlignment) Applier() domain.Applier {
  return e.Apply
}

func (e *ElectricalAlignment) Apply(state domain.State) domain.State {
  // - When attempting to Counterhack, you gain a +20 Base Chance to succeed
  log.Println("applying Electrical Alignment")
  return state
}

var _ domain.Effect = &ElectricalAlignment{}
