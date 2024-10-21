package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type StranglersUnion struct {
  name string
  description string
}

func NewStranglersUnion() *StranglersUnion {
  return &StranglersUnion{
    name: "Strangler's Union",
    description: "When you make an Attack Action with a blackjack, bullwhip or garrote, your foe cannot Dodge or Parry this attack.",
  }
}

func (e *StranglersUnion) Name() string {
  return e.name
}

func (e *StranglersUnion) Description() string {
  return e.description
}

func (e *StranglersUnion) Applier() domain.Applier {
  return e.Apply
}

func (e *StranglersUnion) Apply(state domain.State) domain.State {
  // - When you make an Attack Action with a blackjack, bullwhip or garrote, your foe cannot Dodge or Parry this attack.
  log.Println("applying Strangler's Union")
  return state
}

var _ domain.Effect = &StranglersUnion{}
