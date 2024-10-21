package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type BlessinginDisguise struct {
  name string
  description string
}

func NewBlessinginDisguise() *BlessinginDisguise {
  return &BlessinginDisguise{
    name: "Blessing in Disguise",
    description: "Effect1",
  }
}

func (e *BlessinginDisguise) Name() string {
  return e.name
}

func (e *BlessinginDisguise) Description() string {
  return e.description
}

func (e *BlessinginDisguise) Applier() domain.Applier {
  return e.Apply
}

func (e *BlessinginDisguise) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Blessing in Disguise")
  return state
}

var _ domain.Effect = &BlessinginDisguise{}
