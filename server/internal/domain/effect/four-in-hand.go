package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type FourinHand struct {
  name string
  description string
}

func NewFourinHand() *FourinHand {
  return &FourinHand{
    name: "Four-in-Hand",
    description: "You may flip the results to succeed at Drive Tests. When you succeed, it is always considered a Critical Success. In addition, when using the Movement subtype of Drive, you do not have to add the additional 1 AP cost.",
  }
}

func (e *FourinHand) Name() string {
  return e.name
}

func (e *FourinHand) Description() string {
  return e.description
}

func (e *FourinHand) Applier() domain.Applier {
  return e.Apply
}

func (e *FourinHand) Apply(state domain.State) domain.State {
  // - You may flip the results to succeed at Drive Tests. When you succeed, it is always considered a Critical Success. In addition, when using the Movement subtype of Drive, you do not have to add the additional 1 AP cost.
  log.Println("applying Four-in-Hand")
  return state
}

var _ domain.Effect = &FourinHand{}
