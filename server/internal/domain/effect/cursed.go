package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Cursed struct {
  name string
  description string
}

func NewCursed() *Cursed {
  return &Cursed{
    name: "Cursed",
    description: "Whenever you intend to sacrifice a Fortune Point, roll a 1D6 Chaos Die. If the result is a face ‘6’, you must use two Fortune Points instead of one.",
  }
}

func (e *Cursed) Name() string {
  return e.name
}

func (e *Cursed) Description() string {
  return e.description
}

func (e *Cursed) Applier() domain.Applier {
  return e.Apply
}

func (e *Cursed) Apply(state domain.State) domain.State {
  // - When gossiping or spinning a tale, you gain a +20 Base Chance to Rumor Tests.
  log.Println("applying Cursed")
  return state
}

var _ domain.Effect = &Cursed{}
