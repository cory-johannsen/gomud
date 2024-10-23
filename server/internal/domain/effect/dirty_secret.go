package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type DirtySecret struct {
  name string
  description string
}

func NewDirtySecret() *DirtySecret {
  return &DirtySecret{
    name: "Dirty Secret",
    description: "Whenever the topic that you keep secret comes up, make a Grit roll.  On a Success, you keep your cool.  On a Failure, you Panic.",
  }
}

func (e *DirtySecret) Name() string {
  return e.name
}

func (e *DirtySecret) Description() string {
  return e.description
}

func (e *DirtySecret) Applier() domain.Applier {
  return e.Apply
}

func (e *DirtySecret) Apply(state domain.State) domain.State {
  // - Whenever the topic that you keep secret comes up, make a Grit roll.  On a Success, you keep your cool.  On a Failure, you Panic.
  log.Println("applying Dirty Secret")
  return state
}

var _ domain.Effect = &DirtySecret{}
