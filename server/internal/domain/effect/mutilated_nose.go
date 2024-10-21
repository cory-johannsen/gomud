package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MutilatedNose struct {
  name string
  description string
}

func NewMutilatedNose() *MutilatedNose {
  return &MutilatedNose{
    name: "Mutilated Nose",
    description: "Until fully Recuperated, you cannot smell as you’re in pain. You must undergo a successful surgery or suffer the consequences. Once a Mutilated Nose has undergone a failed surgery, you permanently must flip the results to fail all Skill Tests which rely on smell and taste. If you already have this Drawback, you permanently lose 9% Perception.",
  }
}

func (e *MutilatedNose) Name() string {
  return e.name
}

func (e *MutilatedNose) Description() string {
  return e.description
}

func (e *MutilatedNose) Applier() domain.Applier {
  return e.Apply
}

func (e *MutilatedNose) Apply(state domain.State) domain.State {
  // - Until fully Recuperated, you cannot smell as you’re in pain. You must undergo a successful surgery or suffer the consequences. Once a Mutilated Nose has undergone a failed surgery, you permanently must flip the results to fail all Skill Tests which rely on smell and taste. If you already have this Drawback, you permanently lose 9% Perception.
  log.Println("applying Mutilated Nose")
  return state
}

var _ domain.Effect = &MutilatedNose{}
