package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type LefthandedPath struct {
  name string
  description string
}

func NewLefthandedPath() *LefthandedPath {
  return &LefthandedPath{
    name: "Left-handed Path",
    description: "When you Channel Power, if your Chaos Dice do not result in a Chaos Manifestation, you avoid gaining Corruption as a result.",
  }
}

func (e *LefthandedPath) Name() string {
  return e.name
}

func (e *LefthandedPath) Description() string {
  return e.description
}

func (e *LefthandedPath) Applier() domain.Applier {
  return e.Apply
}

func (e *LefthandedPath) Apply(state domain.State) domain.State {
  // - When you Channel Power, if your Chaos Dice do not result in a Chaos Manifestation, you avoid gaining Corruption as a result.
  log.Println("applying Left-handed Path")
  return state
}

var _ domain.Effect = &LefthandedPath{}
