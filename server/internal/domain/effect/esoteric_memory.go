package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type EsotericMemory struct {
  name string
  description string
}

func NewEsotericMemory() *EsotericMemory {
  return &EsotericMemory{
    name: "Esoteric Memory",
    description: "Effect1",
  }
}

func (e *EsotericMemory) Name() string {
  return e.name
}

func (e *EsotericMemory) Description() string {
  return e.description
}

func (e *EsotericMemory) Applier() domain.Applier {
  return e.Apply
}

func (e *EsotericMemory) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Esoteric Memory")
  return state
}

var _ domain.Effect = &EsotericMemory{}
