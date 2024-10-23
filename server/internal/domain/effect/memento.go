package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Memento struct {
  name string
  description string
}

func NewMemento() *Memento {
  return &Memento{
    name: "Memento",
    description: "When carrying your memento, your Disorder’s effects may be temporarily ignored for 24 hours by spending a Fortune Point,",
  }
}

func (e *Memento) Name() string {
  return e.name
}

func (e *Memento) Description() string {
  return e.description
}

func (e *Memento) Applier() domain.Applier {
  return e.Apply
}

func (e *Memento) Apply(state domain.State) domain.State {
  // - When carrying your memento, your Disorder’s effects may be temporarily ignored for 24 hours by spending a Fortune Point,
  log.Println("applying Memento")
  return state
}

var _ domain.Effect = &Memento{}
