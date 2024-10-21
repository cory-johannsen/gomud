package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type SanguineTemperament struct {
  name string
  description string
}

func NewSanguineTemperament() *SanguineTemperament {
  return &SanguineTemperament{
    name: "Sanguine Temperament",
    description: "You cannot Load or Take Aim without spending an additional Action Point. In addition, whenever you use the Special Action of Wait, you lose 1 Action Point.",
  }
}

func (e *SanguineTemperament) Name() string {
  return e.name
}

func (e *SanguineTemperament) Description() string {
  return e.description
}

func (e *SanguineTemperament) Applier() domain.Applier {
  return e.Apply
}

func (e *SanguineTemperament) Apply(state domain.State) domain.State {
  // - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
  log.Println("applying Sanguine Temperament")
  return state
}

var _ domain.Effect = &SanguineTemperament{}
