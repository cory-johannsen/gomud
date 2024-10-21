package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type EnduringMortality struct {
  name string
  description string
}

func NewEnduringMortality() *EnduringMortality {
  return &EnduringMortality{
    name: "Enduring Mortality",
    description: "Effect1",
  }
}

func (e *EnduringMortality) Name() string {
  return e.name
}

func (e *EnduringMortality) Description() string {
  return e.description
}

func (e *EnduringMortality) Applier() domain.Applier {
  return e.Apply
}

func (e *EnduringMortality) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Enduring Mortality")
  return state
}

var _ domain.Effect = &EnduringMortality{}
