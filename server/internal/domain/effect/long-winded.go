package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Longwinded struct {
  name string
  description string
}

func NewLongwinded() *Longwinded {
  return &Longwinded{
    name: "Long-winded",
    description: "When you suffer physical Peril, reduce your Peril Condition Track by one less step negatively.",
  }
}

func (e *Longwinded) Name() string {
  return e.name
}

func (e *Longwinded) Description() string {
  return e.description
}

func (e *Longwinded) Applier() domain.Applier {
  return e.Apply
}

func (e *Longwinded) Apply(state domain.State) domain.State {
  // - When you suffer physical Peril, reduce your Peril Condition Track by one less step negatively.
  log.Println("applying Long-winded")
  return state
}

var _ domain.Effect = &Longwinded{}
