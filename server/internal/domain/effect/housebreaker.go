package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Housebreaker struct {
  name string
  description string
}

func NewHousebreaker() *Housebreaker {
  return &Housebreaker{
    name: "Housebreaker",
    description: "When trying to open locks, you gain a +20 Base Chance to Skulduggery Tests.",
  }
}

func (e *Housebreaker) Name() string {
  return e.name
}

func (e *Housebreaker) Description() string {
  return e.description
}

func (e *Housebreaker) Applier() domain.Applier {
  return e.Apply
}

func (e *Housebreaker) Apply(state domain.State) domain.State {
  // - When trying to open locks, you gain a +20 Base Chance to Skulduggery Tests.
  log.Println("applying Housebreaker")
  return state
}

var _ domain.Effect = &Housebreaker{}
