package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type BroadBellied struct {
  name string
  description string
}

func NewBroadBellied() *BroadBellied {
  return &BroadBellied{
    name: "Broad Bellied",
    description: "Effect1",
  }
}

func (e *BroadBellied) Name() string {
  return e.name
}

func (e *BroadBellied) Description() string {
  return e.description
}

func (e *BroadBellied) Applier() domain.Applier {
  return e.Apply
}

func (e *BroadBellied) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Broad Bellied")
  return state
}

var _ domain.Effect = &BroadBellied{}
