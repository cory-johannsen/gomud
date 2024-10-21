package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Murdery struct {
  name string
  description string
}

func NewMurdery() *Murdery {
  return &Murdery{
    name: "Murdery",
    description: "'When performing an act of Killing to solve a problem, flip the roll to Success.  Otherwise flip the roll to Failure.'",
  }
}

func (e *Murdery) Name() string {
  return e.name
}

func (e *Murdery) Description() string {
  return e.description
}

func (e *Murdery) Applier() domain.Applier {
  return e.Apply
}

func (e *Murdery) Apply(state domain.State) domain.State {
  // - 'When performing an act of Killing to solve a problem, flip the roll to Success.  Otherwise flip the roll to Failure.'
  log.Println("applying Murdery")
  return state
}

var _ domain.Effect = &Murdery{}
