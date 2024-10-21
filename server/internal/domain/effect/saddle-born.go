package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Saddleborn struct {
  name string
  description string
}

func NewSaddleborn() *Saddleborn {
  return &Saddleborn{
    name: "Saddle-born",
    description: "When fighting on horseback or atop a vehicle like a cart, coach or wagon, you gain a +10 Base Chance to strike with melee and ranged weapons.",
  }
}

func (e *Saddleborn) Name() string {
  return e.name
}

func (e *Saddleborn) Description() string {
  return e.description
}

func (e *Saddleborn) Applier() domain.Applier {
  return e.Apply
}

func (e *Saddleborn) Apply(state domain.State) domain.State {
  // - When fighting on horseback or atop a vehicle like a cart, coach or wagon, you gain a +10 Base Chance to strike with melee and ranged weapons.
  log.Println("applying Saddle-born")
  return state
}

var _ domain.Effect = &Saddleborn{}
