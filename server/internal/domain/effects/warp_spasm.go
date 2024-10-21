package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type WarpSpasm struct {
  Name string
  Description string
}

func (e *WarpSpasm) Apply(state domain.State) domain.State {
  // - Whenever you are Seriously or Grievously Wounded, add a 1D6 Fury Die to Damage you inflict with melee weapons.
  log.Println("applying Warp Spasm")
  return state
}
