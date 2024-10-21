package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type SaltyDog struct {
  name string
  description string
}

func NewSaltyDog() *SaltyDog {
  return &SaltyDog{
    name: "Salty Dog",
    description: "After you make an Attack Action with a melee weapon possessing the Finesse Quality, immediately make an Opportunity Attack with any one-handed ranged weapon on the same Turn. The weapon must already be loaded or in-hand.",
  }
}

func (e *SaltyDog) Name() string {
  return e.name
}

func (e *SaltyDog) Description() string {
  return e.description
}

func (e *SaltyDog) Applier() domain.Applier {
  return e.Apply
}

func (e *SaltyDog) Apply(state domain.State) domain.State {
  // - After you make an Attack Action with a melee weapon possessing the Finesse Quality, immediately make an Opportunity Attack with any one-handed ranged weapon on the same Turn. The weapon must already be loaded or in-hand.
  log.Println("applying Salty Dog")
  return state
}

var _ domain.Effect = &SaltyDog{}
