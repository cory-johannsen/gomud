package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type TunnelVision struct {
  name string
  description string
}

func NewTunnelVision() *TunnelVision {
  return &TunnelVision{
    name: "Tunnel Vision",
    description: "You can see completely in the dark below ground as if it were daylight, providing you are able to use your hands freely.,",
  }
}

func (e *TunnelVision) Name() string {
  return e.name
}

func (e *TunnelVision) Description() string {
  return e.description
}

func (e *TunnelVision) Applier() domain.Applier {
  return e.Apply
}

func (e *TunnelVision) Apply(state domain.State) domain.State {
  // - You can see completely in the dark below ground as if it were daylight, providing you are able to use your hands freely.,
  log.Println("applying Tunnel Vision")
  return state
}

var _ domain.Effect = &TunnelVision{}
