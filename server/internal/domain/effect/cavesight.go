package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Cavesight struct {
  name string
  description string
}

func NewCavesight() *Cavesight {
  return &Cavesight{
    name: "Cavesight",
    description: "Effect1",
  }
}

func (e *Cavesight) Name() string {
  return e.name
}

func (e *Cavesight) Description() string {
  return e.description
}

func (e *Cavesight) Applier() domain.Applier {
  return e.Apply
}

func (e *Cavesight) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Cavesight")
  return state
}

var _ domain.Effect = &Cavesight{}
