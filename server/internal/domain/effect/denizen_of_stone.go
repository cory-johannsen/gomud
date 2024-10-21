package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type DenizenofStone struct {
  name string
  description string
}

func NewDenizenofStone() *DenizenofStone {
  return &DenizenofStone{
    name: "Denizen of Stone",
    description: "Effect1",
  }
}

func (e *DenizenofStone) Name() string {
  return e.name
}

func (e *DenizenofStone) Description() string {
  return e.description
}

func (e *DenizenofStone) Applier() domain.Applier {
  return e.Apply
}

func (e *DenizenofStone) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Denizen of Stone")
  return state
}

var _ domain.Effect = &DenizenofStone{}
