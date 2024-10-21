package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type AimlessDrifter struct {
  name string
  description string
}

func NewAimlessDrifter() *AimlessDrifter {
  return &AimlessDrifter{
    name: "Aimless Drifter",
    description: "When you fail a Folklore or Navigation Test, you may re-roll to generate a better result, but must accept the outcome.",
  }
}

func (e *AimlessDrifter) Name() string {
  return e.name
}

func (e *AimlessDrifter) Description() string {
  return e.description
}

func (e *AimlessDrifter) Applier() domain.Applier {
  return e.Apply
}

func (e *AimlessDrifter) Apply(state domain.State) domain.State {
  // - When you fail a Folklore or Navigation Test, you may re-roll to generate a better result, but must accept the outcome.
  log.Println("applying Aimless Drifter")
  return state
}

var _ domain.Effect = &AimlessDrifter{}
