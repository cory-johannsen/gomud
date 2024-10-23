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
    description: "Reference either your [MB] or [QB] when determining Movement. Additionally, you may substitute Athletics in place of Coordination when Dodging ranged weapons. Finally, you will have a Husky build on the Build table.,",
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
  // - Reference either your [MB] or [QB] when determining Movement. Additionally, you may substitute Athletics in place of Coordination when Dodging ranged weapons. Finally, you will have a Husky build on the Build table.,
  log.Println("applying Physical Prowess")
  return state
}

var _ domain.Effect = &PhysicalProwess{}
