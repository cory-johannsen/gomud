package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type SkullFractureUntil struct {
  name string
  description string
}

func NewSkullFractureUntil() *SkullFractureUntil {
  return &SkullFractureUntil{
    name: "Skull Fracture Until",
    description: "fully Recuperated, you must flip the results to fail all Skill Tests.",
  }
}

func (e *SkullFractureUntil) Name() string {
  return e.name
}

func (e *SkullFractureUntil) Description() string {
  return e.description
}

func (e *SkullFractureUntil) Applier() domain.Applier {
  return e.Apply
}

func (e *SkullFractureUntil) Apply(state domain.State) domain.State {
  // - fully Recuperated, you must flip the results to fail all Skill Tests.
  log.Println("applying Skull Fracture Until")
  return state
}

var _ domain.Effect = &SkullFractureUntil{}
