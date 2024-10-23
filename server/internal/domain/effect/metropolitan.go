package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Metropolitan struct {
  name string
  description string
}

func NewMetropolitan() *Metropolitan {
  return &Metropolitan{
    name: "Metropolitan",
    description: "All trappings you craft (with exception to armor, shields and weapons) are always best-in-class, raising its resale value three times the listed price,",
  }
}

func (e *Metropolitan) Name() string {
  return e.name
}

func (e *Metropolitan) Description() string {
  return e.description
}

func (e *Metropolitan) Applier() domain.Applier {
  return e.Apply
}

func (e *Metropolitan) Apply(state domain.State) domain.State {
  // - All trappings you craft (with exception to armor, shields and weapons) are always best-in-class, raising its resale value three times the listed price,
  log.Println("applying Metropolitan")
  return state
}

var _ domain.Effect = &Metropolitan{}
