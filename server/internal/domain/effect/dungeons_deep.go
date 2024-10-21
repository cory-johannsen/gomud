package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type DungeonsDeep struct {
  name string
  description string
}

func NewDungeonsDeep() *DungeonsDeep {
  return &DungeonsDeep{
    name: "Dungeons Deep",
    description: "Effect1",
  }
}

func (e *DungeonsDeep) Name() string {
  return e.name
}

func (e *DungeonsDeep) Description() string {
  return e.description
}

func (e *DungeonsDeep) Applier() domain.Applier {
  return e.Apply
}

func (e *DungeonsDeep) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Dungeons Deep")
  return state
}

var _ domain.Effect = &DungeonsDeep{}
