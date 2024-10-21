package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type TrueDetective struct {
  Name string
  Description string
}

func (e *TrueDetective) Apply(state domain.State) domain.State {
  // - When Intoxicated or underneath the effects of Deliriants, make a Scrutinize Test. If successful, you can ask for the GM to give you an important clue from your investigations you may not have already thought of or overlooked. This benefit cannot be gained more than once per day.
  log.Println("applying True Detective")
  return state
}
