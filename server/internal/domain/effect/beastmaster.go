package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Beastmaster struct {
  name string
  description string
}

func NewBeastmaster() *Beastmaster {
  return &Beastmaster{
    name: "Beastmaster",
    description: "You can use your Handle Animal Skill to not only tame and train creatures which are classified as Animals, but also those which are classified as Beasts. Whenever you attempt to tame and train Animals and Beasts, you may flip the results to succeed at Handle Animal Tests. When you succeed, it is always considered a Critical Success.",
  }
}

func (e *Beastmaster) Name() string {
  return e.name
}

func (e *Beastmaster) Description() string {
  return e.description
}

func (e *Beastmaster) Applier() domain.Applier {
  return e.Apply
}

func (e *Beastmaster) Apply(state domain.State) domain.State {
  // - You can use your Handle Animal Skill to not only tame and train creatures which are classified as Animals, but also those which are classified as Beasts. Whenever you attempt to tame and train Animals and Beasts, you may flip the results to succeed at Handle Animal Tests. When you succeed, it is always considered a Critical Success.
  log.Println("applying Beastmaster")
  return state
}

var _ domain.Effect = &Beastmaster{}
