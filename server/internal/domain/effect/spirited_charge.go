package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type SpiritedCharge struct {
  name string
  description string
}

func NewSpiritedCharge() *SpiritedCharge {
  return &SpiritedCharge{
    name: "Spirited Charge",
    description: "When you use the Drive or Ride Skills, add +3 to your Movement.",
  }
}

func (e *SpiritedCharge) Name() string {
  return e.name
}

func (e *SpiritedCharge) Description() string {
  return e.description
}

func (e *SpiritedCharge) Applier() domain.Applier {
  return e.Apply
}

func (e *SpiritedCharge) Apply(state domain.State) domain.State {
  // - When you use the Drive or Ride Skills, add +3 to your Movement.
  log.Println("applying Spirited Charge")
  return state
}

var _ domain.Effect = &SpiritedCharge{}
