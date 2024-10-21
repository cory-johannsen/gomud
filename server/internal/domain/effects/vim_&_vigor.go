package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type VimVigor struct {
  Name string
  Description string
}

func (e *VimVigor) Apply(state domain.State) domain.State {
  // - Whenever you Parry a melee weapon, immediately make an Opportunity Attack against that same opponent. You may only make this attack if you are wielding a melee weapon possessing the Finesse Quality.
  log.Println("applying Vim & Vigor")
  return state
}
