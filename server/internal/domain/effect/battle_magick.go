package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type BattleMagick struct {
  name string
  description string
}

func NewBattleMagick() *BattleMagick {
  return &BattleMagick{
    name: "Battle Magick",
    description: "Foes who are made subject to your Magicks must spend one additional Action Point to properly Counterspell. In addition, you penalize your foe’s ability to Resist your Magicks by a -10 Base Chance.",
  }
}

func (e *BattleMagick) Name() string {
  return e.name
}

func (e *BattleMagick) Description() string {
  return e.description
}

func (e *BattleMagick) Applier() domain.Applier {
  return e.Apply
}

func (e *BattleMagick) Apply(state domain.State) domain.State {
  // - Foes who are made subject to your Magicks must spend one additional Action Point to properly Counterspell. In addition, you penalize your foe’s ability to Resist your Magicks by a -10 Base Chance.
  log.Println("applying Battle Magick")
  return state
}

var _ domain.Effect = &BattleMagick{}
