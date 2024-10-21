package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ArtfulDodger struct {
  name string
  description string
}

func NewArtfulDodger() *ArtfulDodger {
  return &ArtfulDodger{
    name: "Artful Dodger",
    description: "You automatically gain every Focus in the Stealth Skill when you enter this Job. This means you may exceed the normal limits for Focuses set by your [IB], but for Stealth only.",
  }
}

func (e *ArtfulDodger) Name() string {
  return e.name
}

func (e *ArtfulDodger) Description() string {
  return e.description
}

func (e *ArtfulDodger) Applier() domain.Applier {
  return e.Apply
}

func (e *ArtfulDodger) Apply(state domain.State) domain.State {
  // - You automatically gain every Focus in the Stealth Skill when you enter this Job. This means you may exceed the normal limits for Focuses set by your [IB], but for Stealth only.
  log.Println("applying Artful Dodger")
  return state
}

var _ domain.Effect = &ArtfulDodger{}
