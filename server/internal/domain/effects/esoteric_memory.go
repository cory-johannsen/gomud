package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type EsotericMemory struct {
  Name string
  Description string
}

func (e *EsotericMemory) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Esoteric Memory")
  return state
}
