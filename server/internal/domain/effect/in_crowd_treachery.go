package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type InCrowdTreachery struct {
  name string
  description string
}

func NewInCrowdTreachery() *InCrowdTreachery {
  return &InCrowdTreachery{
    name: "In Crowd Treachery",
    description: "Your first successful Attack Action against a foe adds a 1D6 Fury Die to your Damage results. This can be used against multiple foes during combat, taking advantage of your traitorous ways!,",
  }
}

func (e *InCrowdTreachery) Name() string {
  return e.name
}

func (e *InCrowdTreachery) Description() string {
  return e.description
}

func (e *InCrowdTreachery) Applier() domain.Applier {
  return e.Apply
}

func (e *InCrowdTreachery) Apply(state domain.State) domain.State {
  // - Your first successful Attack Action against a foe adds a 1D6 Fury Die to your Damage results. This can be used against multiple foes during combat, taking advantage of your traitorous ways!,
  log.Println("applying In Crowd Treachery")
  return state
}

var _ domain.Effect = &InCrowdTreachery{}
