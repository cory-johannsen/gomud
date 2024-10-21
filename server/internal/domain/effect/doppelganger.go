package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Doppelganger struct {
  name string
  description string
}

func NewDoppelganger() *Doppelganger {
  return &Doppelganger{
    name: "Doppelganger",
    description: "When masquerading as someone of a Social Class other than your own, you gain a +20 Base Chance to Disguise Tests.",
  }
}

func (e *Doppelganger) Name() string {
  return e.name
}

func (e *Doppelganger) Description() string {
  return e.description
}

func (e *Doppelganger) Applier() domain.Applier {
  return e.Apply
}

func (e *Doppelganger) Apply(state domain.State) domain.State {
  // - When masquerading as someone of a Social Class other than your own, you gain a +20 Base Chance to Disguise Tests.
  log.Println("applying Doppelganger")
  return state
}

var _ domain.Effect = &Doppelganger{}
