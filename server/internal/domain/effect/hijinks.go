package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Hijinks struct {
  name string
  description string
}

func NewHijinks() *Hijinks {
  return &Hijinks{
    name: "Hijinks",
    description: "Whenever you suffer Damage from a melee weapon, spend a Fortune Point to ignore it entirely.,",
  }
}

func (e *Hijinks) Name() string {
  return e.name
}

func (e *Hijinks) Description() string {
  return e.description
}

func (e *Hijinks) Applier() domain.Applier {
  return e.Apply
}

func (e *Hijinks) Apply(state domain.State) domain.State {
  // - Whenever you suffer Damage from a melee weapon, spend a Fortune Point to ignore it entirely.,
  log.Println("applying Hijinks")
  return state
}

var _ domain.Effect = &Hijinks{}
