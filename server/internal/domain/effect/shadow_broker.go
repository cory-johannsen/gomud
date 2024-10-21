package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ShadowBroker struct {
  name string
  description string
}

func NewShadowBroker() *ShadowBroker {
  return &ShadowBroker{
    name: "Shadow Broker",
    description: "At any time, you may spend 1 Fortune Point to gain a critical piece of information from underworld contacts, rumormongers or other sources.",
  }
}

func (e *ShadowBroker) Name() string {
  return e.name
}

func (e *ShadowBroker) Description() string {
  return e.description
}

func (e *ShadowBroker) Applier() domain.Applier {
  return e.Apply
}

func (e *ShadowBroker) Apply(state domain.State) domain.State {
  // - At any time, you may spend 1 Fortune Point to gain a critical piece of information from underworld contacts, rumormongers or other sources.
  log.Println("applying Shadow Broker")
  return state
}

var _ domain.Effect = &ShadowBroker{}
