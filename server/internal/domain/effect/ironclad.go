package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Ironclad struct {
  name string
  description string
}

func NewIronclad() *Ironclad {
  return &Ironclad{
    name: "Ironclad",
    description: "Effect1",
  }
}

func (e *Ironclad) Name() string {
  return e.name
}

func (e *Ironclad) Description() string {
  return e.description
}

func (e *Ironclad) Applier() domain.Applier {
  return e.Apply
}

func (e *Ironclad) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Ironclad")
  return state
}

var _ domain.Effect = &Ironclad{}
