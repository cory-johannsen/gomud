package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Ammunition struct {
  name string
  description string
}

func NewAmmunition() *Ammunition {
  return &Ammunition{
    name: "Ammunition",
    description: "Weapons of this Quality require ammunition to be used. The weapon will have an amount it holds before it must be reloaded. See Repeating.",
  }
}

func (e *Ammunition) Name() string {
  return e.name
}

func (e *Ammunition) Description() string {
  return e.description
}

func (e *Ammunition) Applier() domain.Applier {
  return e.Apply
}

func (e *Ammunition) Apply(state domain.State) domain.State {
  // - Weapons of this Quality require ammunition to be used. The weapon will have an amount it holds before it must be reloaded. See Repeating.
  log.Println("applying Ammunition")
  return state
}

var _ domain.Effect = &Ammunition{}
