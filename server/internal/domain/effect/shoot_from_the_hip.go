package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ShootFromTheHip struct {
  name string
  description string
}

func NewShootFromTheHip() *ShootFromTheHip {
  return &ShootFromTheHip{
    name: "Shoot From The Hip",
    description: "You can quick draw any one-handed ranged weapon with the Gunpowder Quality for Opportunity Attacks.",
  }
}

func (e *ShootFromTheHip) Name() string {
  return e.name
}

func (e *ShootFromTheHip) Description() string {
  return e.description
}

func (e *ShootFromTheHip) Applier() domain.Applier {
  return e.Apply
}

func (e *ShootFromTheHip) Apply(state domain.State) domain.State {
  // - You can quick draw any one-handed ranged weapon with the Gunpowder Quality for Opportunity Attacks.
  log.Println("applying Shoot From The Hip")
  return state
}

var _ domain.Effect = &ShootFromTheHip{}
