package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Stoneheaded struct {
  name string
  description string
}

func NewStoneheaded() *Stoneheaded {
  return &Stoneheaded{
    name: "Stoneheaded",
    description: "Effect1",
  }
}

func (e *Stoneheaded) Name() string {
  return e.name
}

func (e *Stoneheaded) Description() string {
  return e.description
}

func (e *Stoneheaded) Applier() domain.Applier {
  return e.Apply
}

func (e *Stoneheaded) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Stoneheaded")
  return state
}

var _ domain.Effect = &Stoneheaded{}
