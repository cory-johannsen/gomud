package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type PhysicalProwess struct {
  name string
  description string
}

func NewPhysicalProwess() *PhysicalProwess {
  return &PhysicalProwess{
    name: "Physical Prowess",
    description: "Effect1",
  }
}

func (e *PhysicalProwess) Name() string {
  return e.name
}

func (e *PhysicalProwess) Description() string {
  return e.description
}

func (e *PhysicalProwess) Applier() domain.Applier {
  return e.Apply
}

func (e *PhysicalProwess) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Physical Prowess")
  return state
}

var _ domain.Effect = &PhysicalProwess{}
