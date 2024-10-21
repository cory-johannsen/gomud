package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type JackOfAllTrades struct {
  Name string
  Description string
}

func (e *JackOfAllTrades) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Jack Of All Trades")
  return state
}
