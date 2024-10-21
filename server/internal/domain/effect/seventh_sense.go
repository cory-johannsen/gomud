package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type SeventhSense struct {
  name string
  description string
}

func NewSeventhSense() *SeventhSense {
  return &SeventhSense{
    name: "Seventh Sense",
    description: "Effect1",
  }
}

func (e *SeventhSense) Name() string {
  return e.name
}

func (e *SeventhSense) Description() string {
  return e.description
}

func (e *SeventhSense) Applier() domain.Applier {
  return e.Apply
}

func (e *SeventhSense) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Seventh Sense")
  return state
}

var _ domain.Effect = &SeventhSense{}
