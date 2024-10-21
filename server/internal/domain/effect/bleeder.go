package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Bleeder struct {
  name string
  description string
}

func NewBleeder() *Bleeder {
  return &Bleeder{
    name: "Bleeder",
    description: "Whenever you are treated with the Heal Skill, a caregiver suffers an additional -20 Base Chance, unless they expend an additional bandage during treatment.",
  }
}

func (e *Bleeder) Name() string {
  return e.name
}

func (e *Bleeder) Description() string {
  return e.description
}

func (e *Bleeder) Applier() domain.Applier {
  return e.Apply
}

func (e *Bleeder) Apply(state domain.State) domain.State {
  // - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
  log.Println("applying Bleeder")
  return state
}

var _ domain.Effect = &Bleeder{}
