package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type SilverTongue struct {
  name string
  description string
}

func NewSilverTongue() *SilverTongue {
  return &SilverTongue{
    name: "Silver Tongue",
    description: "When you attempt to persuade those of a different Social Class other than your own, you gain a +20 Base Chance to Charm Tests.",
  }
}

func (e *SilverTongue) Name() string {
  return e.name
}

func (e *SilverTongue) Description() string {
  return e.description
}

func (e *SilverTongue) Applier() domain.Applier {
  return e.Apply
}

func (e *SilverTongue) Apply(state domain.State) domain.State {
  // - When you attempt to persuade those of a different Social Class other than your own, you gain a +20 Base Chance to Charm Tests.
  log.Println("applying Silver Tongue")
  return state
}

var _ domain.Effect = &SilverTongue{}
