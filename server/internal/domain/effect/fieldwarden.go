package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Fieldwarden struct {
  name string
  description string
}

func NewFieldwarden() *Fieldwarden {
  return &Fieldwarden{
    name: "Fieldwarden",
    description: "Effect1",
  }
}

func (e *Fieldwarden) Name() string {
  return e.name
}

func (e *Fieldwarden) Description() string {
  return e.description
}

func (e *Fieldwarden) Applier() domain.Applier {
  return e.Apply
}

func (e *Fieldwarden) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Fieldwarden")
  return state
}

var _ domain.Effect = &Fieldwarden{}
