package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type DontKilltheMessenger struct {
  name string
  description string
}

func NewDontKilltheMessenger() *DontKilltheMessenger {
  return &DontKilltheMessenger{
    name: "Don't Kill the Messenger",
    description: "When using Fellowship-based Skill Tests, you do not suffer any additional penalties due to differences in Social Class.",
  }
}

func (e *DontKilltheMessenger) Name() string {
  return e.name
}

func (e *DontKilltheMessenger) Description() string {
  return e.description
}

func (e *DontKilltheMessenger) Applier() domain.Applier {
  return e.Apply
}

func (e *DontKilltheMessenger) Apply(state domain.State) domain.State {
  // - When using Fellowship-based Skill Tests, you do not suffer any additional penalties due to differences in Social Class.
  log.Println("applying Don't Kill the Messenger")
  return state
}

var _ domain.Effect = &DontKilltheMessenger{}
