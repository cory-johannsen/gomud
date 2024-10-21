package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type OddCouple struct {
  name string
  description string
}

func NewOddCouple() *OddCouple {
  return &OddCouple{
    name: "Odd Couple",
    description: "Effect1",
  }
}

func (e *OddCouple) Name() string {
  return e.name
}

func (e *OddCouple) Description() string {
  return e.description
}

func (e *OddCouple) Applier() domain.Applier {
  return e.Apply
}

func (e *OddCouple) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Odd Couple")
  return state
}

var _ domain.Effect = &OddCouple{}
