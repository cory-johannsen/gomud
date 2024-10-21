package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ThickLining struct {
  name string
  description string
}

func NewThickLining() *ThickLining {
  return &ThickLining{
    name: "Thick Lining",
    description: "Effect1",
  }
}

func (e *ThickLining) Name() string {
  return e.name
}

func (e *ThickLining) Description() string {
  return e.description
}

func (e *ThickLining) Applier() domain.Applier {
  return e.Apply
}

func (e *ThickLining) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Thick Lining")
  return state
}

var _ domain.Effect = &ThickLining{}
