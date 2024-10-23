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
    description: "You are never left Helpless for any reason.,",
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
  // - You are never left Helpless for any reason.,
  log.Println("applying Denizen of Stone")
  return state
}

var _ domain.Effect = &DenizenofStone{}
