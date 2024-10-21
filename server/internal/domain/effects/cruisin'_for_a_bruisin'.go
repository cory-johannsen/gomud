package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type CruisinforaBruisin struct {
  Name string
  Description string
}

func (e *CruisinforaBruisin) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Cruisin' for a Bruisin'")
  return state
}
