package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ThievingWays struct {
  name string
  description string
}

func NewThievingWays() *ThievingWays {
  return &ThievingWays{
    name: "Thieving Ways",
    description: "Effect1",
  }
}

func (e *ThievingWays) Name() string {
  return e.name
}

func (e *ThievingWays) Description() string {
  return e.description
}

func (e *ThievingWays) Applier() domain.Applier {
  return e.Apply
}

func (e *ThievingWays) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Thieving Ways")
  return state
}

var _ domain.Effect = &ThievingWays{}
