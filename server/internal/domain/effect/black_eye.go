package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type BlackEye struct {
  name string
  description string
}

func NewBlackEye() *BlackEye {
  return &BlackEye{
    name: "Black Eye",
    description: "Until fully Recuperated, you must flip the results to fail Skill Tests which rely on vision.",
  }
}

func (e *BlackEye) Name() string {
  return e.name
}

func (e *BlackEye) Description() string {
  return e.description
}

func (e *BlackEye) Applier() domain.Applier {
  return e.Apply
}

func (e *BlackEye) Apply(state domain.State) domain.State {
  // - Until fully Recuperated, you must flip the results to fail Skill Tests which rely on vision.
  log.Println("applying Black Eye")
  return state
}

var _ domain.Effect = &BlackEye{}
