package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ForkedTongue struct {
  name string
  description string
}

func NewForkedTongue() *ForkedTongue {
  return &ForkedTongue{
    name: "Forked Tongue",
    description: "When you attempt to deceive someone of a Social Class other than your own, you gain a +20 Base Chance to Guile Tests.",
  }
}

func (e *ForkedTongue) Name() string {
  return e.name
}

func (e *ForkedTongue) Description() string {
  return e.description
}

func (e *ForkedTongue) Applier() domain.Applier {
  return e.Apply
}

func (e *ForkedTongue) Apply(state domain.State) domain.State {
  // - When you attempt to deceive someone of a Social Class other than your own, you gain a +20 Base Chance to Guile Tests.
  log.Println("applying Forked Tongue")
  return state
}

var _ domain.Effect = &ForkedTongue{}
