package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ThereWillBeBlood struct {
  name string
  description string
}

func NewThereWillBeBlood() *ThereWillBeBlood {
  return &ThereWillBeBlood{
    name: "There Will Be Blood",
    description: "You roll an additional 1D6 Chaos Die to determine if you Injure foes.",
  }
}

func (e *ThereWillBeBlood) Name() string {
  return e.name
}

func (e *ThereWillBeBlood) Description() string {
  return e.description
}

func (e *ThereWillBeBlood) Applier() domain.Applier {
  return e.Apply
}

func (e *ThereWillBeBlood) Apply(state domain.State) domain.State {
  // - You roll an additional 1D6 Chaos Die to determine if you Injure foes.
  log.Println("applying There Will Be Blood")
  return state
}

var _ domain.Effect = &ThereWillBeBlood{}
