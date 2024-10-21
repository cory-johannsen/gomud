package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Metropolitan struct {
  name string
  description string
}

func NewMetropolitan() *Metropolitan {
  return &Metropolitan{
    name: "Metropolitan",
    description: "Effect1",
  }
}

func (e *Metropolitan) Name() string {
  return e.name
}

func (e *Metropolitan) Description() string {
  return e.description
}

func (e *Metropolitan) Applier() domain.Applier {
  return e.Apply
}

func (e *Metropolitan) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Metropolitan")
  return state
}

var _ domain.Effect = &Metropolitan{}
