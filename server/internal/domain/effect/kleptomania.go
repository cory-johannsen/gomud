package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Kleptomania struct {
  name string
  description string
}

func NewKleptomania() *Kleptomania {
  return &Kleptomania{
    name: "Kleptomania",
    description: "Effect1",
  }
}

func (e *Kleptomania) Name() string {
  return e.name
}

func (e *Kleptomania) Description() string {
  return e.description
}

func (e *Kleptomania) Applier() domain.Applier {
  return e.Apply
}

func (e *Kleptomania) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Kleptomania")
  return state
}

var _ domain.Effect = &Kleptomania{}
