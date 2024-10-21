package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type TrueGrit struct {
  name string
  description string
}

func NewTrueGrit() *TrueGrit {
  return &TrueGrit{
    name: "True Grit",
    description: "You are immune to the effect of Knockout! and Stunning Blow.",
  }
}

func (e *TrueGrit) Name() string {
  return e.name
}

func (e *TrueGrit) Description() string {
  return e.description
}

func (e *TrueGrit) Applier() domain.Applier {
  return e.Apply
}

func (e *TrueGrit) Apply(state domain.State) domain.State {
  // - You are immune to the effect of Knockout! and Stunning Blow.
  log.Println("applying True Grit")
  return state
}

var _ domain.Effect = &TrueGrit{}
