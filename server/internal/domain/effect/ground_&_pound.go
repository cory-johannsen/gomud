package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type GroundPound struct {
  name string
  description string
}

func NewGroundPound() *GroundPound {
  return &GroundPound{
    name: "Ground & Pound",
    description: "When a foe successfully Resists your Chokehold, immediately make a bare-handed Opportunity Attack.",
  }
}

func (e *GroundPound) Name() string {
  return e.name
}

func (e *GroundPound) Description() string {
  return e.description
}

func (e *GroundPound) Applier() domain.Applier {
  return e.Apply
}

func (e *GroundPound) Apply(state domain.State) domain.State {
  // - When a foe successfully Resists your Chokehold, immediately make a bare-handed Opportunity Attack.
  log.Println("applying Ground & Pound")
  return state
}

var _ domain.Effect = &GroundPound{}
