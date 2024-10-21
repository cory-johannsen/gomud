package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MountainAmongstMen struct {
  name string
  description string
}

func NewMountainAmongstMen() *MountainAmongstMen {
  return &MountainAmongstMen{
    name: "Mountain Amongst Men",
    description: "Effect1",
  }
}

func (e *MountainAmongstMen) Name() string {
  return e.name
}

func (e *MountainAmongstMen) Description() string {
  return e.description
}

func (e *MountainAmongstMen) Applier() domain.Applier {
  return e.Apply
}

func (e *MountainAmongstMen) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Mountain Amongst Men")
  return state
}

var _ domain.Effect = &MountainAmongstMen{}
