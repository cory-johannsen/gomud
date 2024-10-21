package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Bewitching struct {
  name string
  description string
}

func NewBewitching() *Bewitching {
  return &Bewitching{
    name: "Bewitching",
    description: "Effect1",
  }
}

func (e *Bewitching) Name() string {
  return e.name
}

func (e *Bewitching) Description() string {
  return e.description
}

func (e *Bewitching) Applier() domain.Applier {
  return e.Apply
}

func (e *Bewitching) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Bewitching")
  return state
}

var _ domain.Effect = &Bewitching{}
