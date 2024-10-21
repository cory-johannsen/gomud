package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MountainWarfare struct {
  name string
  description string
}

func NewMountainWarfare() *MountainWarfare {
  return &MountainWarfare{
    name: "Mountain Warfare",
    description: "Effect1",
  }
}

func (e *MountainWarfare) Name() string {
  return e.name
}

func (e *MountainWarfare) Description() string {
  return e.description
}

func (e *MountainWarfare) Applier() domain.Applier {
  return e.Apply
}

func (e *MountainWarfare) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Mountain Warfare")
  return state
}

var _ domain.Effect = &MountainWarfare{}
