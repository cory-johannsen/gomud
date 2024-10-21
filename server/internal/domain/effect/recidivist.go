package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Recidivist struct {
  name string
  description string
}

func NewRecidivist() *Recidivist {
  return &Recidivist{
    name: "Recidivist",
    description: "'When performing an Illegal Act to solve a problem, flip the roll to Success.  Otherwise flip the roll to Failure.'",
  }
}

func (e *Recidivist) Name() string {
  return e.name
}

func (e *Recidivist) Description() string {
  return e.description
}

func (e *Recidivist) Applier() domain.Applier {
  return e.Apply
}

func (e *Recidivist) Apply(state domain.State) domain.State {
  // - 'When performing an Illegal Act to solve a problem, flip the roll to Success.  Otherwise flip the roll to Failure.'
  log.Println("applying Recidivist")
  return state
}

var _ domain.Effect = &Recidivist{}
