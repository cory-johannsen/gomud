package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type TwistedAnkle struct {
  name string
  description string
}

func NewTwistedAnkle() *TwistedAnkle {
  return &TwistedAnkle{
    name: "Twisted Ankle",
    description: "Until fully Recuperated, reduce your Movement by 3.",
  }
}

func (e *TwistedAnkle) Name() string {
  return e.name
}

func (e *TwistedAnkle) Description() string {
  return e.description
}

func (e *TwistedAnkle) Applier() domain.Applier {
  return e.Apply
}

func (e *TwistedAnkle) Apply(state domain.State) domain.State {
  // - Until fully Recuperated, reduce your Movement by 3.
  log.Println("applying Twisted Ankle")
  return state
}

var _ domain.Effect = &TwistedAnkle{}
