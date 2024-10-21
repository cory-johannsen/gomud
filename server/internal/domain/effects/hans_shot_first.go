package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type HansShotFirst struct {
  Name string
  Description string
}

func (e *HansShotFirst) Apply(state domain.State) domain.State {
  // - When combat begins, you gain 3 APs that must be used immediately at the top of the Initiative Order – even if you were Surprised. Once spent, determine your place in the Initiative Ladder and take your Turns normally. If more than two Smugglers are present, the one with the highest [PB] goes first.
  log.Println("applying Hans Shot First")
  return state
}
