package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type LowBlow struct {
  name string
  description string
}

func NewLowBlow() *LowBlow {
  return &LowBlow{
    name: "Low Blow",
    description: "Effect1",
  }
}

func (e *LowBlow) Name() string {
  return e.name
}

func (e *LowBlow) Description() string {
  return e.description
}

func (e *LowBlow) Applier() domain.Applier {
  return e.Apply
}

func (e *LowBlow) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Low Blow")
  return state
}

var _ domain.Effect = &LowBlow{}
