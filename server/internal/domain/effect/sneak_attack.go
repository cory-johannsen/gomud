package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type SneakAttack struct {
  name string
  description string
}

func NewSneakAttack() *SneakAttack {
  return &SneakAttack{
    name: "Sneak Attack",
    description: "Whenever foes are Surprised or outnumbered 6:1 or more, add an additional 1D6 Fury Die to Damage you do against them. You must use a weapon with the Fast Quality. In addition, when using the Movement subtype of Stealth, you do not have to add the additional 1 AP cost.",
  }
}

func (e *SneakAttack) Name() string {
  return e.name
}

func (e *SneakAttack) Description() string {
  return e.description
}

func (e *SneakAttack) Applier() domain.Applier {
  return e.Apply
}

func (e *SneakAttack) Apply(state domain.State) domain.State {
  // - Whenever foes are Surprised or outnumbered 6:1 or more, add an additional 1D6 Fury Die to Damage you do against them. You must use a weapon with the Fast Quality. In addition, when using the Movement subtype of Stealth, you do not have to add the additional 1 AP cost.
  log.Println("applying Sneak Attack")
  return state
}

var _ domain.Effect = &SneakAttack{}
