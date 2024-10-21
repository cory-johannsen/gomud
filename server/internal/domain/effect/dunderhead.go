package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Dunderhead struct {
  name string
  description string
}

func NewDunderhead() *Dunderhead {
  return &Dunderhead{
    name: "Dunderhead",
    description: "Whenever you suffer mental Peril, move one additional step down the Peril Condition Track negatively while suffering 1 Corruption.",
  }
}

func (e *Dunderhead) Name() string {
  return e.name
}

func (e *Dunderhead) Description() string {
  return e.description
}

func (e *Dunderhead) Applier() domain.Applier {
  return e.Apply
}

func (e *Dunderhead) Apply(state domain.State) domain.State {
  // - When gossiping or spinning a tale, you gain a +20 Base Chance to Rumor Tests.
  log.Println("applying Dunderhead")
  return state
}

var _ domain.Effect = &Dunderhead{}
