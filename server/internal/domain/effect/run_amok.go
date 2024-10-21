package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type RunAmok struct {
  name string
  description string
}

func NewRunAmok() *RunAmok {
  return &RunAmok{
    name: "Run Amok",
    description: "When you Charge, you gain a +20 Base Chance to strike with a melee-based Attack Actions and Perilous Stunts.",
  }
}

func (e *RunAmok) Name() string {
  return e.name
}

func (e *RunAmok) Description() string {
  return e.description
}

func (e *RunAmok) Applier() domain.Applier {
  return e.Apply
}

func (e *RunAmok) Apply(state domain.State) domain.State {
  // - When you Charge, you gain a +20 Base Chance to strike with a melee-based Attack Actions and Perilous Stunts.
  log.Println("applying Run Amok")
  return state
}

var _ domain.Effect = &RunAmok{}
