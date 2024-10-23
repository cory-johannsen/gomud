package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type KindredWarband struct {
  name string
  description string
}

func NewKindredWarband() *KindredWarband {
  return &KindredWarband{
    name: "Kindred Warband",
    description: "When fighting alongside others of your Background, you gain a +10 Base Chance to strike with Attack Actions and Perilous Stunts.,",
  }
}

func (e *KindredWarband) Name() string {
  return e.name
}

func (e *KindredWarband) Description() string {
  return e.description
}

func (e *KindredWarband) Applier() domain.Applier {
  return e.Apply
}

func (e *KindredWarband) Apply(state domain.State) domain.State {
  // - When fighting alongside others of your Background, you gain a +10 Base Chance to strike with Attack Actions and Perilous Stunts.,
  log.Println("applying Kindred Warband")
  return state
}

var _ domain.Effect = &KindredWarband{}
