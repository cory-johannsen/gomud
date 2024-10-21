package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type FetteredChaos struct {
  name string
  description string
}

func NewFetteredChaos() *FetteredChaos {
  return &FetteredChaos{
    name: "Fettered Chaos",
    description: "Effect1",
  }
}

func (e *FetteredChaos) Name() string {
  return e.name
}

func (e *FetteredChaos) Description() string {
  return e.description
}

func (e *FetteredChaos) Applier() domain.Applier {
  return e.Apply
}

func (e *FetteredChaos) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Fettered Chaos")
  return state
}

var _ domain.Effect = &FetteredChaos{}
