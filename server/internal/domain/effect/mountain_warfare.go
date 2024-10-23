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
    description: "When wielding any two-handed melee weapon using both hands, any attempt to Parry automatically succeeds (providing you have the Action Points to spend).,",
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
  // - When wielding any two-handed melee weapon using both hands, any attempt to Parry automatically succeeds (providing you have the Action Points to spend).,
  log.Println("applying Mountain Warfare")
  return state
}

var _ domain.Effect = &MountainWarfare{}
