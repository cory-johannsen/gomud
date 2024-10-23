package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type OddCouple struct {
  name string
  description string
}

func NewOddCouple() *OddCouple {
  return &OddCouple{
    name: "Odd Couple",
    description: "When fighting alongside Halfling allies, you gain a +10 Base Chance to strike with Attack Actions and Perilous Stunts.,",
  }
}

func (e *OddCouple) Name() string {
  return e.name
}

func (e *OddCouple) Description() string {
  return e.description
}

func (e *OddCouple) Applier() domain.Applier {
  return e.Apply
}

func (e *OddCouple) Apply(state domain.State) domain.State {
  // - When fighting alongside Halfling allies, you gain a +10 Base Chance to strike with Attack Actions and Perilous Stunts.,
  log.Println("applying Odd Couple")
  return state
}

var _ domain.Effect = &OddCouple{}
