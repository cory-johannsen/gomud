package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ChildrenoftheEarth struct {
  name string
  description string
}

func NewChildrenoftheEarth() *ChildrenoftheEarth {
  return &ChildrenoftheEarth{
    name: "Children of the Earth",
    description: "Effect1",
  }
}

func (e *ChildrenoftheEarth) Name() string {
  return e.name
}

func (e *ChildrenoftheEarth) Description() string {
  return e.description
}

func (e *ChildrenoftheEarth) Applier() domain.Applier {
  return e.Apply
}

func (e *ChildrenoftheEarth) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Children of the Earth")
  return state
}

var _ domain.Effect = &ChildrenoftheEarth{}
