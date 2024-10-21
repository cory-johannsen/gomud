package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type GrimResolve struct {
  name string
  description string
}

func NewGrimResolve() *GrimResolve {
  return &GrimResolve{
    name: "Grim Resolve",
    description: "Effect1",
  }
}

func (e *GrimResolve) Name() string {
  return e.name
}

func (e *GrimResolve) Description() string {
  return e.description
}

func (e *GrimResolve) Applier() domain.Applier {
  return e.Apply
}

func (e *GrimResolve) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Grim Resolve")
  return state
}

var _ domain.Effect = &GrimResolve{}
