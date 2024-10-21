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
    description: "Effect1",
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
  // - Effect1
  log.Println("applying In Crowd Treachery")
  return state
}

var _ domain.Effect = &InCrowdTreachery{}
