package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type PersecutionComplex struct {
  name string
  description string
}

func NewPersecutionComplex() *PersecutionComplex {
  return &PersecutionComplex{
    name: "Persecution Complex",
    description: "You cannot rest to recover from Peril in urban environments, unless you take a dose of laudanum before resting.",
  }
}

func (e *PersecutionComplex) Name() string {
  return e.name
}

func (e *PersecutionComplex) Description() string {
  return e.description
}

func (e *PersecutionComplex) Applier() domain.Applier {
  return e.Apply
}

func (e *PersecutionComplex) Apply(state domain.State) domain.State {
  // - When gossiping or spinning a tale, you gain a +20 Base Chance to Rumor Tests.
  log.Println("applying Persecution Complex")
  return state
}

var _ domain.Effect = &PersecutionComplex{}
