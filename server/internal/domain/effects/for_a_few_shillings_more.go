package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ForAFewShillingsMore struct {
  Name string
  Description string
}

func (e *ForAFewShillingsMore) Apply(state domain.State) domain.State {
  // - In combat, whenever you make your first Attack Action with a ranged weapon, you never miss. In addition, your intended target cannot Dodge, Parry or Resist the attack. You also add an additional 1D6 Fury Die to the same attack.
  log.Println("applying For A Few Shillings More")
  return state
}
