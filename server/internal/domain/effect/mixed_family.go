package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MixedFamily struct {
  name string
  description string
}

func NewMixedFamily() *MixedFamily {
  return &MixedFamily{
    name: "Mixed Family",
    description: "Effect1",
  }
}

func (e *MixedFamily) Name() string {
  return e.name
}

func (e *MixedFamily) Description() string {
  return e.description
}

func (e *MixedFamily) Applier() domain.Applier {
  return e.Apply
}

func (e *MixedFamily) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Mixed Family")
  return state
}

var _ domain.Effect = &MixedFamily{}
