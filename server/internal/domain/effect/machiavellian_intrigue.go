package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MachiavellianIntrigue struct {
  name string
  description string
}

func NewMachiavellianIntrigue() *MachiavellianIntrigue {
  return &MachiavellianIntrigue{
    name: "Machiavellian Intrigue",
    description: "Whenever you fail an Eavesdrop or Scrutinize Test, you may re-roll to generate a better result, but must accept the outcome.",
  }
}

func (e *MachiavellianIntrigue) Name() string {
  return e.name
}

func (e *MachiavellianIntrigue) Description() string {
  return e.description
}

func (e *MachiavellianIntrigue) Applier() domain.Applier {
  return e.Apply
}

func (e *MachiavellianIntrigue) Apply(state domain.State) domain.State {
  // - Whenever you fail an Eavesdrop or Scrutinize Test, you may re-roll to generate a better result, but must accept the outcome.
  log.Println("applying Machiavellian Intrigue")
  return state
}

var _ domain.Effect = &MachiavellianIntrigue{}
