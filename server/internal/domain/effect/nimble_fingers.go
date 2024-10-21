package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type NimbleFingers struct {
  name string
  description string
}

func NewNimbleFingers() *NimbleFingers {
  return &NimbleFingers{
    name: "Nimble Fingers",
    description: "You drastically reduce the time required to use Skulduggery Tests, taking no more than a minute for complex actions (examples include picking locks, pilfering goods, etc.). In addition, you may flip the results to succeed at Skill Tests to steal.",
  }
}

func (e *NimbleFingers) Name() string {
  return e.name
}

func (e *NimbleFingers) Description() string {
  return e.description
}

func (e *NimbleFingers) Applier() domain.Applier {
  return e.Apply
}

func (e *NimbleFingers) Apply(state domain.State) domain.State {
  // - You drastically reduce the time required to use Skulduggery Tests, taking no more than a minute for complex actions (examples include picking locks, pilfering goods, etc.). In addition, you may flip the results to succeed at Skill Tests to steal.
  log.Println("applying Nimble Fingers")
  return state
}

var _ domain.Effect = &NimbleFingers{}
