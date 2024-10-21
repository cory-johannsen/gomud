package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type FortunesMercySerious struct {
  name string
  description string
}

func NewFortunesMercySerious() *FortunesMercySerious {
  return &FortunesMercySerious{
    name: "Fortune’s Mercy! (Serious)",
    description: "Ignore Injury, keep your Fate Point and continue fighting onwards!",
  }
}

func (e *FortunesMercySerious) Name() string {
  return e.name
}

func (e *FortunesMercySerious) Description() string {
  return e.description
}

func (e *FortunesMercySerious) Applier() domain.Applier {
  return e.Apply
}

func (e *FortunesMercySerious) Apply(state domain.State) domain.State {
  // - Ignore Injury, keep your Fate Point and continue fighting onwards!
  log.Println("applying Fortune’s Mercy! (Serious)")
  return state
}

var _ domain.Effect = &FortunesMercySerious{}
