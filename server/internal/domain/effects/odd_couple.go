package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type OddCouple struct {
  Name string
  Description string
}

func (e *OddCouple) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Odd Couple")
  return state
}
