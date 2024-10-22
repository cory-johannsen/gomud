package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Gunpowder struct {
  name string
  description string
}

func NewGunpowder() *Gunpowder {
  return &Gunpowder{
    name: "Gunpowder",
    description: "Weapons of this Quality can be loaded and fired while standing in an Engagement. Furthermore, these weapons cannot be Dodged or Parried. Finally, your Fury Dice explode on a face '1' or '6' when dealing Damage.",
  }
}

func (e *Gunpowder) Name() string {
  return e.name
}

func (e *Gunpowder) Description() string {
  return e.description
}

func (e *Gunpowder) Applier() domain.Applier {
  return e.Apply
}

func (e *Gunpowder) Apply(state domain.State) domain.State {
  // - Weapons of this Quality can be loaded and fired while standing in an Engagement. Furthermore, these weapons cannot be Dodged or Parried. Finally, your Fury Dice explode on a face '1' or '6' when dealing Damage.
  log.Println("applying Gunpowder")
  return state
}

var _ domain.Effect = &Gunpowder{}
