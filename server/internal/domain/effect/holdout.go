package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Holdout struct {
  name string
  description string
}

func NewHoldout() *Holdout {
  return &Holdout{
    name: "Holdout",
    description: "You always succeed at Skulduggery Tests to conceal objects no larger than a knife about your person.",
  }
}

func (e *Holdout) Name() string {
  return e.name
}

func (e *Holdout) Description() string {
  return e.description
}

func (e *Holdout) Applier() domain.Applier {
  return e.Apply
}

func (e *Holdout) Apply(state domain.State) domain.State {
  // - You always succeed at Skulduggery Tests to conceal objects no larger than a knife about your person.
  log.Println("applying Holdout")
  return state
}

var _ domain.Effect = &Holdout{}
