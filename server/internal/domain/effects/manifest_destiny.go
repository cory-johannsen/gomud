package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ManifestDestiny struct {
  Name string
  Description string
}

func (e *ManifestDestiny) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Manifest Destiny")
  return state
}
