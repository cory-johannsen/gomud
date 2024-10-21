package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type TunnelVision struct {
  Name string
  Description string
}

func (e *TunnelVision) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Tunnel Vision")
  return state
}
