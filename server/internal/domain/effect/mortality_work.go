package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MortalityWork struct {
  name string
  description string
}

func NewMortalityWork() *MortalityWork {
  return &MortalityWork{
    name: "Mortality Work",
    description: "When you inflict an Injury with a weapon, your foe also suffers 2D10+your [WB] mental Peril.",
  }
}

func (e *MortalityWork) Name() string {
  return e.name
}

func (e *MortalityWork) Description() string {
  return e.description
}

func (e *MortalityWork) Applier() domain.Applier {
  return e.Apply
}

func (e *MortalityWork) Apply(state domain.State) domain.State {
  // - When you inflict an Injury with a weapon, your foe also suffers 2D10+your [WB] mental Peril.
  log.Println("applying Mortality Work")
  return state
}

var _ domain.Effect = &MortalityWork{}
