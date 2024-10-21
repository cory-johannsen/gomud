package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type FortunesMercy struct {
  name string
  description string
}

func NewFortunesMercy() *FortunesMercy {
  return &FortunesMercy{
    name: "Fortune's Mercy",
    description: "ll",
  }
}

func (e *FortunesMercy) Name() string {
  return e.name
}

func (e *FortunesMercy) Description() string {
  return e.description
}

func (e *FortunesMercy) Applier() domain.Applier {
  return e.Apply
}

func (e *FortunesMercy) Apply(state domain.State) domain.State {
  // null
  log.Println("applying Fortune's Mercy")
  return state
}

var _ domain.Effect = &FortunesMercy{}
