package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MoreWork struct {
  Name string
  Description string
}

func (e *MoreWork) Apply(state domain.State) domain.State {
  // - Adjust your Peril Threshold by +3. In addition, you always recover your Peril Condition Track to Unhindered after resting.
  log.Println("applying More Work?")
  return state
}
