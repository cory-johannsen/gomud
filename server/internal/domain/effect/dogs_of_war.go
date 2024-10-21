package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type DogsofWar struct {
  name string
  description string
}

func NewDogsofWar() *DogsofWar {
  return &DogsofWar{
    name: "Dogs of War",
    description: "You never suffer Serious Injuries.",
  }
}

func (e *DogsofWar) Name() string {
  return e.name
}

func (e *DogsofWar) Description() string {
  return e.description
}

func (e *DogsofWar) Applier() domain.Applier {
  return e.Apply
}

func (e *DogsofWar) Apply(state domain.State) domain.State {
  // - You never suffer Serious Injuries.
  log.Println("applying Dogs of War")
  return state
}

var _ domain.Effect = &DogsofWar{}
