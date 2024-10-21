package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type StrainedGroin struct {
  name string
  description string
}

func NewStrainedGroin() *StrainedGroin {
  return &StrainedGroin{
    name: "Strained Groin",
    description: "Until fully Recuperated, you cannot use any Movement Actions besides Take Cover or Walk in combat.",
  }
}

func (e *StrainedGroin) Name() string {
  return e.name
}

func (e *StrainedGroin) Description() string {
  return e.description
}

func (e *StrainedGroin) Applier() domain.Applier {
  return e.Apply
}

func (e *StrainedGroin) Apply(state domain.State) domain.State {
  // - Until fully Recuperated, you cannot use any Movement Actions besides Take Cover or Walk in combat.
  log.Println("applying Strained Groin")
  return state
}

var _ domain.Effect = &StrainedGroin{}
