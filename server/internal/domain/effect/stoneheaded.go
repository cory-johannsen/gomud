package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Stoneheaded struct {
  name string
  description string
}

func NewStoneheaded() *Stoneheaded {
  return &Stoneheaded{
    name: "Stoneheaded",
    description: "You are immune to Magick which may control your mind and see through illusions without fail.,",
  }
}

func (e *Stoneheaded) Name() string {
  return e.name
}

func (e *Stoneheaded) Description() string {
  return e.description
}

func (e *Stoneheaded) Applier() domain.Applier {
  return e.Apply
}

func (e *Stoneheaded) Apply(state domain.State) domain.State {
  // - You are immune to Magick which may control your mind and see through illusions without fail.,
  log.Println("applying Stoneheaded")
  return state
}

var _ domain.Effect = &Stoneheaded{}
