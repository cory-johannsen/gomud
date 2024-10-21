package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Nighteyes struct {
  name string
  description string
}

func NewNighteyes() *Nighteyes {
  return &Nighteyes{
    name: "Nighteyes",
    description: "Effect1",
  }
}

func (e *Nighteyes) Name() string {
  return e.name
}

func (e *Nighteyes) Description() string {
  return e.description
}

func (e *Nighteyes) Applier() domain.Applier {
  return e.Apply
}

func (e *Nighteyes) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Nighteyes")
  return state
}

var _ domain.Effect = &Nighteyes{}
