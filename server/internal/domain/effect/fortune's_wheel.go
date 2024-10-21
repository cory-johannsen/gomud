package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type FortunesWheel struct {
  name string
  description string
}

func NewFortunesWheel() *FortunesWheel {
  return &FortunesWheel{
    name: "Fortune's Wheel",
    description: "Effect1",
  }
}

func (e *FortunesWheel) Name() string {
  return e.name
}

func (e *FortunesWheel) Description() string {
  return e.description
}

func (e *FortunesWheel) Applier() domain.Applier {
  return e.Apply
}

func (e *FortunesWheel) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Fortune's Wheel")
  return state
}

var _ domain.Effect = &FortunesWheel{}
