package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MachiavellianIntrigue struct {
  Name string
  Description string
}

func (e *MachiavellianIntrigue) Apply(state domain.State) domain.State {
  // - Whenever you fail an Eavesdrop or Scrutinize Test, you may re-roll to generate a better result, but must accept the outcome.
  log.Println("applying Machiavellian Intrigue")
  return state
}
