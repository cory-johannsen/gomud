package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type IncredibleNumeration struct {
  name string
  description string
}

func NewIncredibleNumeration() *IncredibleNumeration {
  return &IncredibleNumeration{
    name: "Incredible Numeration",
    description: "When counting or cheating, you gain a +10 Base Chance to Skill Tests.",
  }
}

func (e *IncredibleNumeration) Name() string {
  return e.name
}

func (e *IncredibleNumeration) Description() string {
  return e.description
}

func (e *IncredibleNumeration) Applier() domain.Applier {
  return e.Apply
}

func (e *IncredibleNumeration) Apply(state domain.State) domain.State {
  // - When counting or cheating, you gain a +10 Base Chance to Skill Tests.
  log.Println("applying Incredible Numeration")
  return state
}

var _ domain.Effect = &IncredibleNumeration{}
