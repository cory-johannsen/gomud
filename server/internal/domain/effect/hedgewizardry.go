package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Hedgewizardry struct {
  name string
  description string
}

func NewHedgewizardry() *Hedgewizardry {
  return &Hedgewizardry{
    name: "Hedgewizardry",
    description: "When casting Generalist Magick – at your option – you may automatically succeed at the Incantation Test, but move one step down the Peril Condition Track negatively. You also understand how to use the Ritual of Magick Circle.",
  }
}

func (e *Hedgewizardry) Name() string {
  return e.name
}

func (e *Hedgewizardry) Description() string {
  return e.description
}

func (e *Hedgewizardry) Applier() domain.Applier {
  return e.Apply
}

func (e *Hedgewizardry) Apply(state domain.State) domain.State {
  // - When casting Generalist Magick – at your option – you may automatically succeed at the Incantation Test, but move one step down the Peril Condition Track negatively. You also understand how to use the Ritual of Magick Circle.
  log.Println("applying Hedgewizardry")
  return state
}

var _ domain.Effect = &Hedgewizardry{}
