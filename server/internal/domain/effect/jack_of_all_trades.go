package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type JackOfAllTrades struct {
  name string
  description string
}

func NewJackOfAllTrades() *JackOfAllTrades {
  return &JackOfAllTrades{
    name: "Jack Of All Trades",
    description: "Effect1",
  }
}

func (e *JackOfAllTrades) Name() string {
  return e.name
}

func (e *JackOfAllTrades) Description() string {
  return e.description
}

func (e *JackOfAllTrades) Applier() domain.Applier {
  return e.Apply
}

func (e *JackOfAllTrades) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Jack Of All Trades")
  return state
}

var _ domain.Effect = &JackOfAllTrades{}
