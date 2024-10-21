package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type SecretSigns struct {
  name string
  description string
}

func NewSecretSigns() *SecretSigns {
  return &SecretSigns{
    name: "Secret Signs",
    description: "When you attempt to understand secret symbols, hand gestures or signs left behind by others as a warning, you gain a +10 Base Chance to Skill Tests.",
  }
}

func (e *SecretSigns) Name() string {
  return e.name
}

func (e *SecretSigns) Description() string {
  return e.description
}

func (e *SecretSigns) Applier() domain.Applier {
  return e.Apply
}

func (e *SecretSigns) Apply(state domain.State) domain.State {
  // - When you attempt to understand secret symbols, hand gestures or signs left behind by others as a warning, you gain a +10 Base Chance to Skill Tests.
  log.Println("applying Secret Signs")
  return state
}

var _ domain.Effect = &SecretSigns{}
