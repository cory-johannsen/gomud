package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type NostrumRemedium struct {
  Name string
  Description string
}

func (e *NostrumRemedium) Apply(state domain.State) domain.State {
  // - You may flip the results to succeed at Alchemy Tests. When you succeed, it is always considered a Critical Success. In addition, you never suffer Peril as a result of failed or Critically Failed Alchemy Tests.
  log.Println("applying Nostrum Remedium")
  return state
}
