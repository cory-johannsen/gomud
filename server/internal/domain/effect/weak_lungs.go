package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type WeakLungs struct {
  name string
  description string
}

func NewWeakLungs() *WeakLungs {
  return &WeakLungs{
    name: "Weak Lungs",
    description: "Whenever you suffer physical Peril, move one additional step down the Peril Condition Track negatively while suffering 1 Corruption.",
  }
}

func (e *WeakLungs) Name() string {
  return e.name
}

func (e *WeakLungs) Description() string {
  return e.description
}

func (e *WeakLungs) Applier() domain.Applier {
  return e.Apply
}

func (e *WeakLungs) Apply(state domain.State) domain.State {
  // - When gossiping or spinning a tale, you gain a +20 Base Chance to Rumor Tests.
  log.Println("applying Weak Lungs")
  return state
}

var _ domain.Effect = &WeakLungs{}
