package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type NobleSavage struct {
  name string
  description string
}

func NewNobleSavage() *NobleSavage {
  return &NobleSavage{
    name: "Noble Savage",
    description: "You never suffer physical Peril as a result of failed Toughness Tests,",
  }
}

func (e *NobleSavage) Name() string {
  return e.name
}

func (e *NobleSavage) Description() string {
  return e.description
}

func (e *NobleSavage) Applier() domain.Applier {
  return e.Apply
}

func (e *NobleSavage) Apply(state domain.State) domain.State {
  // - You never suffer physical Peril as a result of failed Toughness Tests,
  log.Println("applying Noble Savage")
  return state
}

var _ domain.Effect = &NobleSavage{}
