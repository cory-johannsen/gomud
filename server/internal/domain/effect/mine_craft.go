package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MineCraft struct {
  name string
  description string
}

func NewMineCraft() *MineCraft {
  return &MineCraft{
    name: "Mine Craft",
    description: "When beneath the earth, you gain a +20 Base Chance to Navigation Tests. In addition, you always know true north below ground, even in total darkness.",
  }
}

func (e *MineCraft) Name() string {
  return e.name
}

func (e *MineCraft) Description() string {
  return e.description
}

func (e *MineCraft) Applier() domain.Applier {
  return e.Apply
}

func (e *MineCraft) Apply(state domain.State) domain.State {
  // - When beneath the earth, you gain a +20 Base Chance to Navigation Tests. In addition, you always know true north below ground, even in total darkness.
  log.Println("applying Mine Craft")
  return state
}

var _ domain.Effect = &MineCraft{}
