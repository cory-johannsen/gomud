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
    description: "You can never be caught in a Chokehold and may flip the results to succeed at Coordination Tests to slip through bonds and other situations which may detain you from moving.,",
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
  // - You can never be caught in a Chokehold and may flip the results to succeed at Coordination Tests to slip through bonds and other situations which may detain you from moving.,
  log.Println("applying Escape Artist")
  return state
}

var _ domain.Effect = &EscapeArtist{}
