package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type VimVigor struct {
  name string
  description string
}

func NewVimVigor() *VimVigor {
  return &VimVigor{
    name: "Vim & Vigor",
    description: "Whenever you Parry a melee weapon, immediately make an Opportunity Attack against that same opponent. You may only make this attack if you are wielding a melee weapon possessing the Finesse Quality.",
  }
}

func (e *VimVigor) Name() string {
  return e.name
}

func (e *VimVigor) Description() string {
  return e.description
}

func (e *VimVigor) Applier() domain.Applier {
  return e.Apply
}

func (e *VimVigor) Apply(state domain.State) domain.State {
  // - Whenever you Parry a melee weapon, immediately make an Opportunity Attack against that same opponent. You may only make this attack if you are wielding a melee weapon possessing the Finesse Quality.
  log.Println("applying Vim & Vigor")
  return state
}

var _ domain.Effect = &VimVigor{}
