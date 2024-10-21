package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type FracturedLarynx struct {
  name string
  description string
}

func NewFracturedLarynx() *FracturedLarynx {
  return &FracturedLarynx{
    name: "Fractured Larynx",
    description: "Until fully Recuperated, you must succeed at a Scrutinize Test to speak.",
  }
}

func (e *FracturedLarynx) Name() string {
  return e.name
}

func (e *FracturedLarynx) Description() string {
  return e.description
}

func (e *FracturedLarynx) Applier() domain.Applier {
  return e.Apply
}

func (e *FracturedLarynx) Apply(state domain.State) domain.State {
  // - Until fully Recuperated, you must succeed at a Scrutinize Test to speak.
  log.Println("applying Fractured Larynx")
  return state
}

var _ domain.Effect = &FracturedLarynx{}
