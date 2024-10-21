package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type CheapShot struct {
  name string
  description string
}

func NewCheapShot() *CheapShot {
  return &CheapShot{
    name: "Cheap Shot",
    description: "When a foe successfully Parries your Melee Attack, immediately make a bare-handed Opportunity Attack.",
  }
}

func (e *CheapShot) Name() string {
  return e.name
}

func (e *CheapShot) Description() string {
  return e.description
}

func (e *CheapShot) Applier() domain.Applier {
  return e.Apply
}

func (e *CheapShot) Apply(state domain.State) domain.State {
  // - When a foe successfully Parries your Melee Attack, immediately make a bare-handed Opportunity Attack.
  log.Println("applying Cheap Shot")
  return state
}

var _ domain.Effect = &CheapShot{}
