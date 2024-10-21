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
    description: "Effect1",
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
  // - Effect1
  log.Println("applying Tunnel Vision")
  return state
}

var _ domain.Effect = &TunnelVision{}
