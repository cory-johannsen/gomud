package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Slamdance struct {
  name string
  description string
}

func NewSlamdance() *Slamdance {
  return &Slamdance{
    name: "Slamdance",
    description: "Effect1",
  }
}

func (e *Slamdance) Name() string {
  return e.name
}

func (e *Slamdance) Description() string {
  return e.description
}

func (e *Slamdance) Applier() domain.Applier {
  return e.Apply
}

func (e *Slamdance) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Slamdance")
  return state
}

var _ domain.Effect = &Slamdance{}
