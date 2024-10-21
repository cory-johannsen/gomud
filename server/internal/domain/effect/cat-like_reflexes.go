package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type CatlikeReflexes struct {
  name string
  description string
}

func NewCatlikeReflexes() *CatlikeReflexes {
  return &CatlikeReflexes{
    name: "Cat-like Reflexes",
    description: "Effect1",
  }
}

func (e *CatlikeReflexes) Name() string {
  return e.name
}

func (e *CatlikeReflexes) Description() string {
  return e.description
}

func (e *CatlikeReflexes) Applier() domain.Applier {
  return e.Apply
}

func (e *CatlikeReflexes) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Cat-like Reflexes")
  return state
}

var _ domain.Effect = &CatlikeReflexes{}
