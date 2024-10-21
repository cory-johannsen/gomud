package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Mastermind struct {
  name string
  description string
}

func NewMastermind() *Mastermind {
  return &Mastermind{
    name: "Mastermind",
    description: "When you fail a Folklore or Intimidate Test, you may re-roll to generate a better result, but must accept the outcome.",
  }
}

func (e *Mastermind) Name() string {
  return e.name
}

func (e *Mastermind) Description() string {
  return e.description
}

func (e *Mastermind) Applier() domain.Applier {
  return e.Apply
}

func (e *Mastermind) Apply(state domain.State) domain.State {
  // - When you fail a Folklore or Intimidate Test, you may re-roll to generate a better result, but must accept the outcome.
  log.Println("applying Mastermind")
  return state
}

var _ domain.Effect = &Mastermind{}
