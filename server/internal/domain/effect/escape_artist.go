package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type EscapeArtist struct {
  name string
  description string
}

func NewEscapeArtist() *EscapeArtist {
  return &EscapeArtist{
    name: "Escape Artist",
    description: "Effect1",
  }
}

func (e *EscapeArtist) Name() string {
  return e.name
}

func (e *EscapeArtist) Description() string {
  return e.description
}

func (e *EscapeArtist) Applier() domain.Applier {
  return e.Apply
}

func (e *EscapeArtist) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Escape Artist")
  return state
}

var _ domain.Effect = &EscapeArtist{}
