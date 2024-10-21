package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type NervesofSteel struct {
  name string
  description string
}

func NewNervesofSteel() *NervesofSteel {
  return &NervesofSteel{
    name: "Nerves of Steel",
    description: "When resting in unsafe places, your Peril Condition Track moves all steps positively to Unhindered.",
  }
}

func (e *NervesofSteel) Name() string {
  return e.name
}

func (e *NervesofSteel) Description() string {
  return e.description
}

func (e *NervesofSteel) Applier() domain.Applier {
  return e.Apply
}

func (e *NervesofSteel) Apply(state domain.State) domain.State {
  // - When resting in unsafe places, your Peril Condition Track moves all steps positively to Unhindered.
  log.Println("applying Nerves of Steel")
  return state
}

var _ domain.Effect = &NervesofSteel{}
