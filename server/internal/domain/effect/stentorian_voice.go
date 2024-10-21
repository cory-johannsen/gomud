package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type StentorianVoice struct {
  name string
  description string
}

func NewStentorianVoice() *StentorianVoice {
  return &StentorianVoice{
    name: "Stentorian Voice",
    description: "Effect1",
  }
}

func (e *StentorianVoice) Name() string {
  return e.name
}

func (e *StentorianVoice) Description() string {
  return e.description
}

func (e *StentorianVoice) Applier() domain.Applier {
  return e.Apply
}

func (e *StentorianVoice) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Stentorian Voice")
  return state
}

var _ domain.Effect = &StentorianVoice{}
