package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type PaythePiper struct {
  name string
  description string
}

func NewPaythePiper() *PaythePiper {
  return &PaythePiper{
    name: "Pay the Piper",
    description: "Whenever you fail a Bargain or Intimidate Test, you may re-roll to generate a better result, but must accept the outcome.",
  }
}

func (e *PaythePiper) Name() string {
  return e.name
}

func (e *PaythePiper) Description() string {
  return e.description
}

func (e *PaythePiper) Applier() domain.Applier {
  return e.Apply
}

func (e *PaythePiper) Apply(state domain.State) domain.State {
  // - Whenever you fail a Bargain or Intimidate Test, you may re-roll to generate a better result, but must accept the outcome.
  log.Println("applying Pay the Piper")
  return state
}

var _ domain.Effect = &PaythePiper{}
