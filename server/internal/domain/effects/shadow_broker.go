package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ShadowBroker struct {
  Name string
  Description string
}

func (e *ShadowBroker) Apply(state domain.State) domain.State {
  // - At any time, you may spend 1 Fortune Point to gain a critical piece of information from underworld contacts, rumormongers or other sources.
  log.Println("applying Shadow Broker")
  return state
}
