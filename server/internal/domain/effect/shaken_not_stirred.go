package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ShakenNotStirred struct {
  name string
  description string
}

func NewShakenNotStirred() *ShakenNotStirred {
  return &ShakenNotStirred{
    name: "Shaken Not Stirred",
    description: "When you fail a Charm or Eavesdrop Test, you may re-roll to generate a better result, but must accept the outcome.",
  }
}

func (e *ShakenNotStirred) Name() string {
  return e.name
}

func (e *ShakenNotStirred) Description() string {
  return e.description
}

func (e *ShakenNotStirred) Applier() domain.Applier {
  return e.Apply
}

func (e *ShakenNotStirred) Apply(state domain.State) domain.State {
  // - When you fail a Charm or Eavesdrop Test, you may re-roll to generate a better result, but must accept the outcome.
  log.Println("applying Shaken Not Stirred")
  return state
}

var _ domain.Effect = &ShakenNotStirred{}
