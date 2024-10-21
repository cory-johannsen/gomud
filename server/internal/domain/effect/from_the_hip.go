package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type FromTheHip struct {
  name string
  description string
}

func NewFromTheHip() *FromTheHip {
  return &FromTheHip{
    name: "From The Hip",
    description: "From The Hip Effect",
  }
}

func (e *FromTheHip) Name() string {
  return e.name
}

func (e *FromTheHip) Description() string {
  return e.description
}

func (e *FromTheHip) Applier() domain.Applier {
  return e.Apply
}

func (e *FromTheHip) Apply(state domain.State) domain.State {
  // - From The Hip Effect
  log.Println("applying From The Hip")
  return state
}

var _ domain.Effect = &FromTheHip{}
