package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type WarpSpasm struct {
  name string
  description string
}

func NewWarpSpasm() *WarpSpasm {
  return &WarpSpasm{
    name: "Warp Spasm",
    description: "Whenever you are Seriously or Grievously Wounded, add a 1D6 Fury Die to Damage you inflict with melee weapons.",
  }
}

func (e *WarpSpasm) Name() string {
  return e.name
}

func (e *WarpSpasm) Description() string {
  return e.description
}

func (e *WarpSpasm) Applier() domain.Applier {
  return e.Apply
}

func (e *WarpSpasm) Apply(state domain.State) domain.State {
  // - Whenever you are Seriously or Grievously Wounded, add a 1D6 Fury Die to Damage you inflict with melee weapons.
  log.Println("applying Warp Spasm")
  return state
}

var _ domain.Effect = &WarpSpasm{}
