package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type NaturesOwn struct {
  name string
  description string
}

func NewNaturesOwn() *NaturesOwn {
  return &NaturesOwn{
    name: "Nature's Own",
    description: "You leave no trace of your passing in rural areas or above ground, unless discovered by Magick or a Critically Succeeded Awareness Test.,",
  }
}

func (e *NaturesOwn) Name() string {
  return e.name
}

func (e *NaturesOwn) Description() string {
  return e.description
}

func (e *NaturesOwn) Applier() domain.Applier {
  return e.Apply
}

func (e *NaturesOwn) Apply(state domain.State) domain.State {
  // - You leave no trace of your passing in rural areas or above ground, unless discovered by Magick or a Critically Succeeded Awareness Test.,
  log.Println("applying Nature's Own")
  return state
}

var _ domain.Effect = &NaturesOwn{}
