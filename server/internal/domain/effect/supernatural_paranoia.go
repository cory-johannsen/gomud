package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type SupernaturalParanoia struct {
  name string
  description string
}

func NewSupernaturalParanoia() *SupernaturalParanoia {
  return &SupernaturalParanoia{
    name: "Supernatural Paranoia",
    description: "When your Chaos Ranks are higher than your Order Ranks, add 3 to your Initiative.",
  }
}

func (e *SupernaturalParanoia) Name() string {
  return e.name
}

func (e *SupernaturalParanoia) Description() string {
  return e.description
}

func (e *SupernaturalParanoia) Applier() domain.Applier {
  return e.Apply
}

func (e *SupernaturalParanoia) Apply(state domain.State) domain.State {
  // - When your Chaos Ranks are higher than your Order Ranks, add 3 to your Initiative.
  log.Println("applying Supernatural Paranoia")
  return state
}

var _ domain.Effect = &SupernaturalParanoia{}
