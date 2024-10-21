package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type SeveredArtery struct {
  name string
  description string
}

func NewSeveredArtery() *SeveredArtery {
  return &SeveredArtery{
    name: "Severed Artery",
    description: "An arterial spray of blood marks your doom; you are instantly Slain!",
  }
}

func (e *SeveredArtery) Name() string {
  return e.name
}

func (e *SeveredArtery) Description() string {
  return e.description
}

func (e *SeveredArtery) Applier() domain.Applier {
  return e.Apply
}

func (e *SeveredArtery) Apply(state domain.State) domain.State {
  // - An arterial spray of blood marks your doom; you are instantly Slain!
  log.Println("applying Severed Artery")
  return state
}

var _ domain.Effect = &SeveredArtery{}
