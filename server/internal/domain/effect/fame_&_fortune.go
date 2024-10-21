package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type FameFortune struct {
  name string
  description string
}

func NewFameFortune() *FameFortune {
  return &FameFortune{
    name: "Fame & Fortune",
    description: "When making any Skill Test, you never suffer the ill-effect of Critical Failures, instead treating it as a failed Skill Test.",
  }
}

func (e *FameFortune) Name() string {
  return e.name
}

func (e *FameFortune) Description() string {
  return e.description
}

func (e *FameFortune) Applier() domain.Applier {
  return e.Apply
}

func (e *FameFortune) Apply(state domain.State) domain.State {
  // - When making any Skill Test, you never suffer the ill-effect of Critical Failures, instead treating it as a failed Skill Test.
  log.Println("applying Fame & Fortune")
  return state
}

var _ domain.Effect = &FameFortune{}
