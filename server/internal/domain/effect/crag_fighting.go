package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type CragFighting struct {
  name string
  description string
}

func NewCragFighting() *CragFighting {
  return &CragFighting{
    name: "Crag Fighting",
    description: "Effect1",
  }
}

func (e *CragFighting) Name() string {
  return e.name
}

func (e *CragFighting) Description() string {
  return e.description
}

func (e *CragFighting) Applier() domain.Applier {
  return e.Apply
}

func (e *CragFighting) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Crag Fighting")
  return state
}

var _ domain.Effect = &CragFighting{}
