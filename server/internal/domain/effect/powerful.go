package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Powerful struct {
  name string
  description string
}

func NewPowerful() *Powerful {
  return &Powerful{
    name: "Powerful",
    description: "Immediately after striking an Engaged foe, weapons of this Quality force a foe to Resist with Toughness or be shoved out of the Engagement.",
  }
}

func (e *Powerful) Name() string {
  return e.name
}

func (e *Powerful) Description() string {
  return e.description
}

func (e *Powerful) Applier() domain.Applier {
  return e.Apply
}

func (e *Powerful) Apply(state domain.State) domain.State {
  // - Immediately after striking an Engaged foe, weapons of this Quality force a foe to Resist with Toughness or be shoved out of the Engagement.
  log.Println("applying Powerful")
  return state
}

var _ domain.Effect = &Powerful{}
