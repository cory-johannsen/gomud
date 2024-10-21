package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type BruisedRibs struct {
  name string
  description string
}

func NewBruisedRibs() *BruisedRibs {
  return &BruisedRibs{
    name: "Bruised Ribs",
    description: "Until fully Recuperated, reduce your Encumbrance Limit by 3.",
  }
}

func (e *BruisedRibs) Name() string {
  return e.name
}

func (e *BruisedRibs) Description() string {
  return e.description
}

func (e *BruisedRibs) Applier() domain.Applier {
  return e.Apply
}

func (e *BruisedRibs) Apply(state domain.State) domain.State {
  // - Until fully Recuperated, reduce your Encumbrance Limit by 3.
  log.Println("applying Bruised Ribs")
  return state
}

var _ domain.Effect = &BruisedRibs{}
