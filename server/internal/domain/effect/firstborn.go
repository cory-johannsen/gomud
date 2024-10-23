package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Firstborn struct {
  name string
  description string
}

func NewFirstborn() *Firstborn {
  return &Firstborn{
    name: "Firstborn",
    description: "Whenever learning new Magicks, you always succeed at your Incantation Test. Additionally, you never have to sacrifice Reward Points to finalize the process,",
  }
}

func (e *Firstborn) Name() string {
  return e.name
}

func (e *Firstborn) Description() string {
  return e.description
}

func (e *Firstborn) Applier() domain.Applier {
  return e.Apply
}

func (e *Firstborn) Apply(state domain.State) domain.State {
  // - Whenever learning new Magicks, you always succeed at your Incantation Test. Additionally, you never have to sacrifice Reward Points to finalize the process,
  log.Println("applying Firstborn")
  return state
}

var _ domain.Effect = &Firstborn{}
