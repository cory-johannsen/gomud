package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MelancholicTemperament struct {
  name string
  description string
}

func NewMelancholicTemperament() *MelancholicTemperament {
  return &MelancholicTemperament{
    name: "Melancholic Temperament",
    description: "Whenever you use Smelling Salts, you must use two instead of one. Using an additional dose of Smelling Salts in this instance has no negative effect.",
  }
}

func (e *MelancholicTemperament) Name() string {
  return e.name
}

func (e *MelancholicTemperament) Description() string {
  return e.description
}

func (e *MelancholicTemperament) Applier() domain.Applier {
  return e.Apply
}

func (e *MelancholicTemperament) Apply(state domain.State) domain.State {
  // - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
  log.Println("applying Melancholic Temperament")
  return state
}

var _ domain.Effect = &MelancholicTemperament{}
