package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Grudgebearer struct {
  name string
  description string
}

func NewGrudgebearer() *Grudgebearer {
  return &Grudgebearer{
    name: "Grudgebearer",
    description: "Effect1",
  }
}

func (e *Grudgebearer) Name() string {
  return e.name
}

func (e *Grudgebearer) Description() string {
  return e.description
}

func (e *Grudgebearer) Applier() domain.Applier {
  return e.Apply
}

func (e *Grudgebearer) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Grudgebearer")
  return state
}

var _ domain.Effect = &Grudgebearer{}
