package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type BrokenRib struct {
  name string
  description string
}

func NewBrokenRib() *BrokenRib {
  return &BrokenRib{
    name: "Broken Rib",
    description: "Your armor gains the Ruined! Quality. Until fully Recuperated, you cannot add Skill Ranks to Combat, Brawn or Agility-based Skills.",
  }
}

func (e *BrokenRib) Name() string {
  return e.name
}

func (e *BrokenRib) Description() string {
  return e.description
}

func (e *BrokenRib) Applier() domain.Applier {
  return e.Apply
}

func (e *BrokenRib) Apply(state domain.State) domain.State {
  // - Your armor gains the Ruined! Quality. Until fully Recuperated, you cannot add Skill Ranks to Combat, Brawn or Agility-based Skills.
  log.Println("applying Broken Rib")
  return state
}

var _ domain.Effect = &BrokenRib{}
