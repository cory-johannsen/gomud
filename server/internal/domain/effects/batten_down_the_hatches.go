package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type BattenDowntheHatches struct {
  Name string
  Description string
}

func (e *BattenDowntheHatches) Apply(state domain.State) domain.State {
  // - You may flip the results to succeed at Pilot Tests. When you succeed, it is always considered a Critical Success. In addition, when using the Movement subtype of Swim, you do not have to add the additional 1 Action Point cost.
  log.Println("applying Batten Down the Hatches")
  return state
}
