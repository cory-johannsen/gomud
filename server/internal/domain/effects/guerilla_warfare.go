package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type GuerillaWarfare struct {
  Name string
  Description string
}

func (e *GuerillaWarfare) Apply(state domain.State) domain.State {
  // - You never provoke Opportunity Attacks with Movement Actions or any other action you take. In addition, any Movement Action a foe takes while Engaged with you immediately provokes an Opportunity Attack from you.
  log.Println("applying Guerilla Warfare")
  return state
}
