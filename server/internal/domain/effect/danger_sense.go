package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type DangerSense struct {
  name string
  description string
}

func NewDangerSense() *DangerSense {
  return &DangerSense{
    name: "Danger Sense",
    description: "Effect1",
  }
}

func (e *DangerSense) Name() string {
  return e.name
}

func (e *DangerSense) Description() string {
  return e.description
}

func (e *DangerSense) Applier() domain.Applier {
  return e.Apply
}

func (e *DangerSense) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Danger Sense")
  return state
}

var _ domain.Effect = &DangerSense{}
