package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Punishing struct {
  name string
  description string
}

func NewPunishing() *Punishing {
  return &Punishing{
    name: "Punishing",
    description: "Immediately after striking a foe, weapons of this Quality may add a 1D6 Fury Die to Total Damage in exchange for spending 1 additional Action Point on this Turn. You’ll learn more about Action Points and Fury Dice in Chapter 8",
  }
}

func (e *Punishing) Name() string {
  return e.name
}

func (e *Punishing) Description() string {
  return e.description
}

func (e *Punishing) Applier() domain.Applier {
  return e.Apply
}

func (e *Punishing) Apply(state domain.State) domain.State {
  // - Immediately after striking a foe, weapons of this Quality may add a 1D6 Fury Die to Total Damage in exchange for spending 1 additional Action Point on this Turn. You’ll learn more about Action Points and Fury Dice in Chapter 8
  log.Println("applying Punishing")
  return state
}

var _ domain.Effect = &Punishing{}
