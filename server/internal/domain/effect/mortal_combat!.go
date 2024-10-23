package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MortalCombat struct {
  name string
  description string
}

func NewMortalCombat() *MortalCombat {
  return &MortalCombat{
    name: "Mortal Combat!",
    description: "You ignore the Pummeling and Weak Qualities when fighting with Brawling type of weapons. In addition, you may refer to [MB] or [BB] when inflicting Damage with this type of weapon, whichever is more favorable.",
  }
}

func (e *MortalCombat) Name() string {
  return e.name
}

func (e *MortalCombat) Description() string {
  return e.description
}

func (e *MortalCombat) Applier() domain.Applier {
  return e.Apply
}

func (e *MortalCombat) Apply(state domain.State) domain.State {
  // - You ignore the Pummeling and Weak Qualities when fighting with Brawling type of weapons. In addition, you may refer to [MB] or [BB] when inflicting Damage with this type of weapon, whichever is more favorable.
  log.Println("applying Mortal Combat!")
  return state
}

var _ domain.Effect = &MortalCombat{}
