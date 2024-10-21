package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type BanterandJibe struct {
  Name string
  Description string
}

func (e *BanterandJibe) Apply(state domain.State) domain.State {
  // - You may flip the results to succeed at Coordination Tests. When you succeed, it is always considered a Critical Success. In addition, you always succeed at Coordination Tests to perform acrobatics.
  log.Println("applying Banter and Jibe")
  return state
}
