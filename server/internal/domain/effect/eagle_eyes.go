package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type EagleEyes struct {
  name string
  description string
}

func NewEagleEyes() *EagleEyes {
  return &EagleEyes{
    name: "Eagle Eyes",
    description: "You do not suffer additional hardship when firing ranged weapons at Medium Distance, instead treating it as Short Distance",
  }
}

func (e *EagleEyes) Name() string {
  return e.name
}

func (e *EagleEyes) Description() string {
  return e.description
}

func (e *EagleEyes) Applier() domain.Applier {
  return e.Apply
}

func (e *EagleEyes) Apply(state domain.State) domain.State {
  // - You do not suffer additional hardship when firing ranged weapons at Medium Distance, instead treating it as Short Distance
  log.Println("applying Eagle Eyes")
  return state
}

var _ domain.Effect = &EagleEyes{}
