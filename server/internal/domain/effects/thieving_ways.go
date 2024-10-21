package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ThievingWays struct {
  Name string
  Description string
}

func (e *ThievingWays) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Thieving Ways")
  return state
}
