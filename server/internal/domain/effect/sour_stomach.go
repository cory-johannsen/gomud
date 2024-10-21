package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type SourStomach struct {
  name string
  description string
}

func NewSourStomach() *SourStomach {
  return &SourStomach{
    name: "Sour Stomach",
    description: "After taking a dose of Laudanum, a Delirient or consuming a substance the GM sees as being hard on your stomach, you cannot recover to Unhindered on the Peril Condition Track (only to Imperiled) for the next 24-hours.",
  }
}

func (e *SourStomach) Name() string {
  return e.name
}

func (e *SourStomach) Description() string {
  return e.description
}

func (e *SourStomach) Applier() domain.Applier {
  return e.Apply
}

func (e *SourStomach) Apply(state domain.State) domain.State {
  // - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
  log.Println("applying Sour Stomach")
  return state
}

var _ domain.Effect = &SourStomach{}
