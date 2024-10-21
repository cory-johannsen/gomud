package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type SituationalAwareness struct {
  name string
  description string
}

func NewSituationalAwareness() *SituationalAwareness {
  return &SituationalAwareness{
    name: "Situational Awareness",
    description: "When you fail an Awareness or Stealth Test, you may re-roll to generate a better result, but must accept the outcome.",
  }
}

func (e *SituationalAwareness) Name() string {
  return e.name
}

func (e *SituationalAwareness) Description() string {
  return e.description
}

func (e *SituationalAwareness) Applier() domain.Applier {
  return e.Apply
}

func (e *SituationalAwareness) Apply(state domain.State) domain.State {
  // - When you fail an Awareness or Stealth Test, you may re-roll to generate a better result, but must accept the outcome.
  log.Println("applying Situational Awareness")
  return state
}

var _ domain.Effect = &SituationalAwareness{}
