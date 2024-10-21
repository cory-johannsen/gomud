package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type CastIronStomach struct {
  name string
  description string
}

func NewCastIronStomach() *CastIronStomach {
  return &CastIronStomach{
    name: "Cast Iron Stomach",
    description: "Effect1",
  }
}

func (e *CastIronStomach) Name() string {
  return e.name
}

func (e *CastIronStomach) Description() string {
  return e.description
}

func (e *CastIronStomach) Applier() domain.Applier {
  return e.Apply
}

func (e *CastIronStomach) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Cast Iron Stomach")
  return state
}

var _ domain.Effect = &CastIronStomach{}
