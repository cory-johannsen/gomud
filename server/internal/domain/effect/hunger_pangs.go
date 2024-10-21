package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type HungerPangs struct {
  name string
  description string
}

func NewHungerPangs() *HungerPangs {
  return &HungerPangs{
    name: "Hunger Pangs",
    description: "Effect1",
  }
}

func (e *HungerPangs) Name() string {
  return e.name
}

func (e *HungerPangs) Description() string {
  return e.description
}

func (e *HungerPangs) Applier() domain.Applier {
  return e.Apply
}

func (e *HungerPangs) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Hunger Pangs")
  return state
}

var _ domain.Effect = &HungerPangs{}
