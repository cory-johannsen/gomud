package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type TownGossip struct {
  Name string
  Description string
}

func (e *TownGossip) Apply(state domain.State) domain.State {
  // - When you fail an Eavesdrop or Rumor Test, you may re-roll to generate a better result, but must accept the outcome.
  log.Println("applying Town Gossip")
  return state
}
