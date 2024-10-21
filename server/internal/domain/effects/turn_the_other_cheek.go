package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type TurntheOtherCheek struct {
  Name string
  Description string
}

func (e *TurntheOtherCheek) Apply(state domain.State) domain.State {
  // - You can read, write and speak the language of any creature that is classified as a Humanoid (including player Ancestries). In addition, whenever you suffer Corruption, decrease the number you gain by 1 (to a minimum of one). This means that if you suffer 3 Corruption, you gain 2 instead.
  log.Println("applying Turn the Other Cheek")
  return state
}
